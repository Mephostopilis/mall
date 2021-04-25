package service

import (
	"context"

	pb "edu/api/logic/grpc"
)

func (s *LogicService) NodesWeighted(ctx context.Context, req *pb.NodesWeightedReq) (reply *pb.NodesWeightedReply, err error) {
	// TODO: 远程ip
	// remoteIP := c.RemoteIP()
	remoteIP := ""
	_ = s.uc.NodesWeighted(ctx, req.Platform, remoteIP)
	reply = &pb.NodesWeightedReply{}
	return
}

func (s *LogicService) NodesInstances(ctx context.Context, req *pb.NodesInstancesReq) (reply *pb.NodesInstancesReply, err error) {
	return
}
