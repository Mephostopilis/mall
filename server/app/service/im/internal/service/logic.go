package service

import (
	"context"

	pb "edu/api/logic/grpc"
	"edu/service/im/internal/biz"
)

type LogicService struct {
	pb.UnimplementedLogicServer
	uc *biz.Logic
}

func NewLogicService(uc *biz.Logic) *LogicService {
	return &LogicService{
		uc: uc,
	}
}

func (s *LogicService) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingReply, error) {
	return &pb.PingReply{}, nil
}
func (s *LogicService) Close(ctx context.Context, req *pb.CloseReq) (*pb.CloseReply, error) {
	return &pb.CloseReply{}, nil
}
func (s *LogicService) Connect(ctx context.Context, req *pb.ConnectReq) (*pb.ConnectReply, error) {
	return &pb.ConnectReply{}, nil
}
func (s *LogicService) Disconnect(ctx context.Context, req *pb.DisconnectReq) (*pb.DisconnectReply, error) {
	return &pb.DisconnectReply{}, nil
}
func (s *LogicService) Heartbeat(ctx context.Context, req *pb.HeartbeatReq) (*pb.HeartbeatReply, error) {
	return &pb.HeartbeatReply{}, nil
}
func (s *LogicService) RenewOnline(ctx context.Context, req *pb.OnlineReq) (*pb.OnlineReply, error) {
	return &pb.OnlineReply{}, nil
}
func (s *LogicService) Receive(ctx context.Context, req *pb.ReceiveReq) (*pb.ReceiveReply, error) {
	return &pb.ReceiveReply{}, nil
}
func (s *LogicService) Nodes(ctx context.Context, req *pb.NodesReq) (*pb.NodesReply, error) {
	return &pb.NodesReply{}, nil
}
