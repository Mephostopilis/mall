package biz

import (
	"context"
	"fmt"
	"net/url"
	"sync/atomic"
	"time"

	comet "edu/api/comet/grpc"
	"edu/job/handleim/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// grpc options
	grpcKeepAliveTime    = time.Duration(10) * time.Second
	grpcKeepAliveTimeout = time.Duration(3) * time.Second
	grpcBackoffMaxDelay  = time.Duration(3) * time.Second
	grpcMaxSendMsgSize   = 1 << 24
	grpcMaxCallMsgSize   = 1 << 24
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
)

// func newCometClient(addr string) (comet.CometClient, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
// 	defer cancel()
// 	conn, err := grpc.DialContext(ctx, addr,
// 		[]grpc.DialOption{
// 			grpc.WithInsecure(),
// 			grpc.WithInitialWindowSize(grpcInitialWindowSize),
// 			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
// 			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
// 			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
// 			grpc.WithBackoffMaxDelay(grpcBackoffMaxDelay),
// 			grpc.WithKeepaliveParams(keepalive.ClientParameters{
// 				Time:                grpcKeepAliveTime,
// 				Timeout:             grpcKeepAliveTimeout,
// 				PermitWithoutStream: true,
// 			}),
// 		}...,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return comet.NewCometClient(conn), err
// }

// Comet is a comet.
type Comet struct {
	serverID      string
	client        comet.CometClient
	pushChan      []chan *comet.PushMsgReq
	roomChan      []chan *comet.BroadcastRoomReq
	broadcastChan chan *comet.BroadcastReq
	pushChanNum   uint64
	roomChanNum   uint64
	routineSize   uint64

	ctx    context.Context
	cancel context.CancelFunc
	log    *log.Helper
}

// NewComet new a comet.
func NewComet(regc *conf.Reg, c *conf.Server_Comet) (cmt *Comet, err error) {
	cmt = &Comet{
		serverID:      regc.Instance.Hostname,
		pushChan:      make([]chan *comet.PushMsgReq, c.RoutineSize),
		roomChan:      make([]chan *comet.BroadcastRoomReq, c.RoutineSize),
		broadcastChan: make(chan *comet.BroadcastReq, c.RoutineSize),
		routineSize:   uint64(c.RoutineSize),
	}
	var grpcAddr string
	for _, addrs := range regc.Instance.Addrs {
		u, err := url.Parse(addrs)
		if err == nil && u.Scheme == "grpc" {
			grpcAddr = u.Host
		}
	}
	if grpcAddr == "" {
		return nil, fmt.Errorf("invalid grpc address:%v", regc.Instance.Addrs)
	}
	if cmt.client, err = comet.NewClient(context.Background()); err != nil {
		return nil, err
	}
	cmt.ctx, cmt.cancel = context.WithCancel(context.Background())

	for i := int32(0); i < c.RoutineSize; i++ {
		cmt.pushChan[i] = make(chan *comet.PushMsgReq, c.RoutineChan)
		cmt.roomChan[i] = make(chan *comet.BroadcastRoomReq, c.RoutineChan)
		go cmt.process(cmt.pushChan[i], cmt.roomChan[i], cmt.broadcastChan)
	}
	return cmt, nil
}

// Push push a user message.
func (c *Comet) Push(arg *comet.PushMsgReq) (err error) {
	idx := atomic.AddUint64(&c.pushChanNum, 1) % c.routineSize
	c.pushChan[idx] <- arg
	return
}

// BroadcastRoom broadcast a room message.
func (c *Comet) BroadcastRoom(arg *comet.BroadcastRoomReq) (err error) {
	idx := atomic.AddUint64(&c.roomChanNum, 1) % c.routineSize
	c.roomChan[idx] <- arg
	return
}

// Broadcast broadcast a message.
func (c *Comet) Broadcast(arg *comet.BroadcastReq) (err error) {
	c.broadcastChan <- arg
	return
}

func (c *Comet) process(pushChan chan *comet.PushMsgReq, roomChan chan *comet.BroadcastRoomReq, broadcastChan chan *comet.BroadcastReq) {
	for {
		select {
		case broadcastArg := <-broadcastChan:
			_, err := c.client.Broadcast(context.Background(), &comet.BroadcastReq{
				Proto:   broadcastArg.Proto,
				ProtoOp: broadcastArg.ProtoOp,
				Speed:   broadcastArg.Speed,
			})
			if err != nil {
				c.log.Errorf("c.client.Broadcast(%s, reply) serverId:%s error(%v)", broadcastArg, c.serverID, err)
			}
		case roomArg := <-roomChan:
			_, err := c.client.BroadcastRoom(context.Background(), &comet.BroadcastRoomReq{
				RoomID: roomArg.RoomID,
				Proto:  roomArg.Proto,
			})
			if err != nil {
				c.log.Errorf("c.client.BroadcastRoom(%s, reply) serverId:%s error(%v)", roomArg, c.serverID, err)
			}
		case pushArg := <-pushChan:
			_, err := c.client.PushMsg(context.Background(), &comet.PushMsgReq{
				Keys:    pushArg.Keys,
				Proto:   pushArg.Proto,
				ProtoOp: pushArg.ProtoOp,
			})
			if err != nil {
				c.log.Errorf("c.client.PushMsg(%s, reply) serverId:%s error(%v)", pushArg, c.serverID, err)
			}
		case <-c.ctx.Done():
			return
		}
	}
}

// Close close the resouces.
func (c *Comet) Close() (err error) {
	finish := make(chan bool)
	go func() {
		for {
			n := len(c.broadcastChan)
			for _, ch := range c.pushChan {
				n += len(ch)
			}
			for _, ch := range c.roomChan {
				n += len(ch)
			}
			if n == 0 {
				finish <- true
				return
			}
			time.Sleep(time.Second)
		}
	}()
	select {
	case <-finish:
		c.log.Infof("close comet finish")
	case <-time.After(5 * time.Second):
		err = fmt.Errorf("close comet(server:%s push:%d room:%d broadcast:%d) timeout", c.serverID, len(c.pushChan), len(c.roomChan), len(c.broadcastChan))
	}
	c.cancel()
	return
}
