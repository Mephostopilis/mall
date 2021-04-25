package biz

import (
	"context"
	"math/rand"
	"runtime"
	"time"

	logic "edu/api/logic/grpc"
	"edu/comet/internal/conf"
	"edu/pkg/net/ip"

	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/uuid"
	"github.com/zhenjl/cityhash"
)

const (
	minServerHeartbeat = time.Minute * 10
	maxServerHeartbeat = time.Minute * 30
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
	grpcMaxSendMsgSize        = 1 << 24
	grpcMaxCallMsgSize        = 1 << 24
	grpcKeepAliveTime         = time.Second * 10
	grpcKeepAliveTimeout      = time.Second * 3
	grpcBackoffMaxDelay       = time.Second * 3
)

// Server is comet server.
type Server struct {
	Debug     bool
	c         *conf.Server
	uuid      string
	round     *Round    // accept round store
	buckets   []*Bucket // subkey bucket
	bucketIdx uint32

	serverID  string
	rpcClient logic.LogicClient
	log       *log.Helper
}

// NewServer returns a new Server.
func NewServer(id uuid.UUID, csrv *conf.Service, c *conf.Server, logger log.Logger, r *etcd.Registry) (*Server, error) {
	ins := &registry.ServiceInstance{
		ID:        id.String(),
		Name:      csrv.Name,
		Version:   csrv.Version,
		Metadata:  map[string]string{},
		Endpoints: []string{c.Grpc.Address},
	}
	if err := r.Register(context.Background(), ins); err != nil {
		return nil, err
	}
	logicClient, err := logic.NewClient(context.Background(), grpc.WithDiscovery(r))
	if err != nil {
		return nil, err
	}
	s := &Server{
		c:         c,
		uuid:      id.String(),
		round:     NewRound(c),
		rpcClient: logicClient,
	}
	// init bucket
	s.buckets = make([]*Bucket, c.Bucket.Size)
	s.bucketIdx = uint32(c.Bucket.Size)
	for i := int32(0); i < c.Bucket.Size; i++ {
		s.buckets[i] = NewBucket(c.Bucket)
	}
	s.serverID = ip.InternalIP()
	go s.onlineproc()

	if err := InitWhitelist(c.Whitelist, logger); err != nil {
		panic(err)
	}
	if err := InitTCP(logger, s, c.Tcp.Bind, runtime.NumCPU()); err != nil {
		panic(err)
	}
	if err := InitWebsocket(logger, s, c.Websocket.Bind, runtime.NumCPU()); err != nil {
		panic(err)
	}
	// if conf.Conf.Websocket.TLSOpen {
	// 	if err := comet.InitWebsocketWithTLS(srv, conf.Conf.Websocket.TLSBind, conf.Conf.Websocket.CertFile, conf.Conf.Websocket.PrivateFile, runtime.NumCPU()); err != nil {
	// 		panic(err)
	// 	}
	// }

	return s, nil
}

// Buckets return all buckets.
func (s *Server) Buckets() []*Bucket {
	return s.buckets
}

// Bucket get the bucket by subkey.
func (s *Server) Bucket(subKey string) *Bucket {
	idx := cityhash.CityHash32([]byte(subKey), uint32(len(subKey))) % s.bucketIdx
	if s.Debug {
		s.log.Infof("%s hit channel bucket index: %d use cityhash", subKey, idx)
	}
	return s.buckets[idx]
}

// RandServerHearbeat rand server heartbeat.
func (s *Server) RandServerHearbeat() time.Duration {
	return (minServerHeartbeat + time.Duration(rand.Int63n(int64(maxServerHeartbeat-minServerHeartbeat))))
}

// Close close the server.
func (s *Server) Close() (err error) {
	return
}

func (s *Server) onlineproc() {
	for {
		var (
			allRoomsCount map[string]int32
			err           error
		)
		roomCount := make(map[string]int32)
		for _, bucket := range s.buckets {
			for roomID, count := range bucket.RoomsCount() {
				roomCount[roomID] += count
			}
		}
		if allRoomsCount, err = s.RenewOnline(context.Background(), s.serverID, roomCount); err != nil {
			time.Sleep(time.Second)
			continue
		}
		for _, bucket := range s.buckets {
			bucket.UpRoomsCount(allRoomsCount)
		}
		time.Sleep(time.Second * 10)
	}
}
