package service

import (
	"context"

	pb "edu/api/logic/grpc"
)

func (s *LogicService) OnlineRoom(ctx context.Context, req *pb.OnlineRoomReq) (reply *pb.OnlineRoomReply, err error) {
	res, err := s.uc.OnlineRoom(ctx, req.Type, req.Rooms)
	if err != nil {
		return
	}
	reply = &pb.OnlineRoomReply{
		Data: res,
	}
	return
}

func (s *LogicService) OnlineTop(ctx context.Context, req *pb.OnlineTopReq) (reply *pb.OnlineTopReply, err error) {
	res, err := s.uc.OnlineTop(ctx, req.Type, int(req.Limit))
	if err != nil {
		return
	}
	list := make([]*pb.Top, 0)
	for i := 0; i < len(res); i++ {
		it := res[i]
		list = append(list, &pb.Top{
			RoomId: it.RoomID,
			Count:  it.Count,
		})
	}
	reply = &pb.OnlineTopReply{
		List: list,
	}
	return
}

func (s *LogicService) OnlineTotal(ctx context.Context, req *pb.OnlineTotalReq) (reply *pb.OnlineTotalReply, err error) {
	ipCount, connCount := s.uc.OnlineTotal(ctx)
	reply = &pb.OnlineTotalReply{
		IpCount:   ipCount,
		ConnCount: connCount,
	}
	return
}
