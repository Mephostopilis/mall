package service

import (
	"context"
	"time"

	pb "edu/api/comet/grpc"
	"edu/comet/internal/biz"
	"edu/comet/internal/errors"
)

type CometService struct {
	pb.UnimplementedCometServer
	uc *biz.Server
}

func NewCometService(s *biz.Server) *CometService {
	return &CometService{
		uc: s,
	}
}

func (s *CometService) Ping(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
func (s *CometService) Close(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
func (s *CometService) PushMsg(ctx context.Context, req *pb.PushMsgReq) (reply *pb.PushMsgReply, err error) {
	if len(req.Keys) == 0 || req.Proto == nil {
		return nil, errors.ErrPushMsgArg
	}
	for _, key := range req.Keys {
		if channel := s.uc.Bucket(key).Channel(key); channel != nil {
			if !channel.NeedPush(req.ProtoOp) {
				continue
			}
			if err = channel.Push(req.Proto); err != nil {
				return
			}
		}
	}
	return &pb.PushMsgReply{}, nil
}

func (s *CometService) Broadcast(ctx context.Context, req *pb.BroadcastReq) (*pb.BroadcastReply, error) {
	if req.Proto == nil {
		return nil, errors.ErrBroadCastArg
	}
	// TODO use broadcast queue
	go func() {
		for _, bucket := range s.uc.Buckets() {
			bucket.Broadcast(req.GetProto(), req.ProtoOp)
			if req.Speed > 0 {
				t := bucket.ChannelCount() / int(req.Speed)
				time.Sleep(time.Duration(t) * time.Second)
			}
		}
	}()
	return &pb.BroadcastReply{}, nil
}

func (s *CometService) BroadcastRoom(ctx context.Context, req *pb.BroadcastRoomReq) (*pb.BroadcastRoomReply, error) {
	if req.Proto == nil || req.RoomID == "" {
		return nil, errors.ErrBroadCastRoomArg
	}
	for _, bucket := range s.uc.Buckets() {
		bucket.BroadcastRoom(req)
	}
	return &pb.BroadcastRoomReply{}, nil
}

func (s *CometService) Rooms(ctx context.Context, req *pb.RoomsReq) (*pb.RoomsReply, error) {
	var (
		roomIds = make(map[string]bool)
	)
	for _, bucket := range s.uc.Buckets() {
		for roomID := range bucket.Rooms() {
			roomIds[roomID] = true
		}
	}
	return &pb.RoomsReply{Rooms: roomIds}, nil
}
