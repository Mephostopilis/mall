package biz

import (
	"context"
	"strconv"
	"time"

	"edu/service/im/internal/conf"
	"edu/service/im/internal/dao"
	"edu/service/im/internal/model"

	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

const (
	_onlineTick     = time.Second * 10
	_onlineDeadline = time.Minute * 5
)

const AppID = "discovery:///comet"

// Logic struct
type Logic struct {
	c   *conf.App
	dis *etcd.Registry
	dao dao.Dao
	// online
	totalIPs   int64
	totalConns int64
	roomCount  map[string]int32
	// load balancer
	nodes        []*registry.ServiceInstance
	loadBalancer *LoadBalancer
	regions      map[string]string // province -> region
	log          *log.Helper
}

// New init
func New(c *conf.App, logger log.Logger, r *etcd.Registry, d dao.Dao) (l *Logic) {
	l = &Logic{
		c:            c,
		dao:          d,
		dis:          r,
		loadBalancer: NewLoadBalancer(),
		regions:      make(map[string]string),
		log:          log.NewHelper(log.With(logger, "module", "biz/logic")),
	}
	l.initRegions()
	l.initNodes()
	_ = l.loadOnline()
	go l.onlineproc()
	return l
}

// Ping ping resources is ok.
func (l *Logic) Ping(c context.Context) (err error) {
	return l.dao.Ping(c)
}

// Close close resources.
func (l *Logic) Close() {
	l.dao.Close()
}

func (l *Logic) initRegions() {
	// for region, ps := range l.c.Regions {
	// 	for _, province := range ps {
	// 		l.regions[province] = region
	// 	}
	// }
}

func (l *Logic) initNodes() {
	nodes, err := l.dis.GetService(context.Background(), AppID)
	if err != nil {
		panic("discovery watch failed")
	}
	for i := 0; i < len(nodes); i++ {
		l.newNodes(nodes[i])
	}
	go func() {
		for {
			w, err := l.dis.Watch(context.Background(), AppID)
			if err != nil {
				continue
			}
			nodes, err := w.Next()
			if err != nil {
				continue
			}
			for i := 0; i < len(nodes); i++ {
				l.newNodes(nodes[i])
			}
		}
	}()
}

func (l *Logic) newNodes(ins *registry.ServiceInstance) {
	var (
		totalConns int64
		totalIPs   int64
		allIns     []*registry.ServiceInstance
	)

	if ins.Metadata == nil {
		l.log.Errorf("node instance metadata is empty(%+v)", ins)
		return
	}
	offline, err := strconv.ParseBool(ins.Metadata[model.MetaOffline])
	if err != nil || offline {
		l.log.Warnf("strconv.ParseBool(offline:%t) error(%v)", offline, err)
		return
	}
	conns, err := strconv.ParseInt(ins.Metadata[model.MetaConnCount], 10, 32)
	if err != nil {
		l.log.Errorf("strconv.ParseInt(conns:%d) error(%v)", conns, err)
		return
	}
	ips, err := strconv.ParseInt(ins.Metadata[model.MetaIPCount], 10, 32)
	if err != nil {
		l.log.Errorf("strconv.ParseInt(ips:%d) error(%v)", ips, err)
		return
	}
	totalConns += conns
	totalIPs += ips
	allIns = append(allIns, ins)

	l.totalConns = totalConns
	l.totalIPs = totalIPs
	l.nodes = allIns
	l.loadBalancer.Update(allIns)
}

func (l *Logic) onlineproc() {
	for {
		time.Sleep(_onlineTick)
		if err := l.loadOnline(); err != nil {
			l.log.Errorf("onlineproc error(%v)", err)
		}
	}
}

func (l *Logic) loadOnline() (err error) {
	var (
		roomCount = make(map[string]int32)
	)
	for _, server := range l.nodes {
		var online *model.Online
		online, err = l.dao.ServerOnline(context.Background(), server.Endpoints[0])
		if err != nil {
			return
		}
		if time.Since(time.Unix(online.Updated, 0)) > _onlineDeadline {
			_ = l.dao.DelServerOnline(context.Background(), server.Endpoints[0])
			continue
		}
		for roomID, count := range online.RoomCount {
			roomCount[roomID] += count
		}
	}
	l.roomCount = roomCount
	return
}
