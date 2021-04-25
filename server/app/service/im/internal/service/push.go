package service

import (
	"context"

	pb "edu/api/logic/grpc"
)

// TODO: 推送key不对
func (s *LogicService) PushKeys(ctx context.Context, req *pb.PushKeysReq) (reply *pb.PushKeysReply, err error) {
	if err = s.uc.PushKeys(ctx, req.Operation, req.Keys, []byte("")); err != nil {
		return
	}
	reply = &pb.PushKeysReply{}
	return
}

func (s *LogicService) PushMids(ctx context.Context, req *pb.PushMidsReq) (reply *pb.PushMidsReply, err error) {
	if err = s.uc.PushMids(ctx, req.Operation, req.Mids, []byte("")); err != nil {
		return
	}
	reply = &pb.PushMidsReply{}
	return
}

func (s *LogicService) PushRoom(ctx context.Context, req *pb.PushRoomReq) (reply *pb.PushRoomReply, err error) {
	if err = s.uc.PushRoom(ctx, req.Operation, req.Type, req.Room, []byte("")); err != nil {
		return
	}
	reply = &pb.PushRoomReply{}
	return
}

func (s *LogicService) PushAll(ctx context.Context, req *pb.PushAllReq) (reply *pb.PushAllReply, err error) {
	if err = s.uc.PushAll(ctx, req.Operation, req.Speed, []byte("")); err != nil {
		return
	}
	return &pb.PushAllReply{}, nil
}
