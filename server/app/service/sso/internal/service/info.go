package service

import (
	"context"

	pb "edu/api/sso/v1"
)

func (s *ApiService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (reply *pb.GetUserInfoResp, err error) {
	// u, err := s.dao.GetUser(ctx, res.Mid)
	// if err != nil {
	// 	return
	// }
	// reply = &pb.GetUserInfoResp{
	// 	U: &pb.User{
	// 		ID:          u.ID,
	// 		Vip:         u.Vip,
	// 		VipMode:     u.VipMode,
	// 		Authority:   "admin",
	// 		Avatar:      "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png",
	// 		Email:       "123@163.com",
	// 		Signature:   "海纳百川，有容乃大",
	// 		Title:       "海纳百川，有容乃大",
	// 		NotifyCount: 12,
	// 		Country:     "China",
	// 		Address:     "西湖区工专路 77 号",
	// 		Phone:       "0752-268888888",
	// 	},
	// }
	// if u.Vip > 0 {
	// 	now := time.Now()
	// 	if now.Before(u.VipDeadline.Time()) {
	// 		ti := u.VipDeadline.Time()
	// 		du := ti.Sub(now)
	// 		reply.U.VipRemaining = du.String()
	// 	} else {
	// 		u.Vip = 0
	// 		s.dao.PutUser(ctx, u)
	// 	}
	// }
	// return
	return
}
