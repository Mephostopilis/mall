package grpc

import (
	"context"
	"time"

	grpctransport "github.com/go-kratos/kratos/v2/transport/grpc"
)

// AppID .
const appID = "discovery:///im"

// NewClient new grpc client
func NewClient(ctx context.Context, opts ...grpctransport.ClientOption) (LogicClient, error) {
	t := make([]grpctransport.ClientOption, 0)
	t = append(t, grpctransport.WithEndpoint(appID))
	t = append(t, grpctransport.WithTimeout(time.Minute))
	t = append(t, opts...)
	cc, err := grpctransport.DialInsecure(ctx, t...)
	if err != nil {
		return nil, err
	}
	// cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	// cc, err := client.Dial(context.Background(), fmt.Sprintf("etcd://default/%s", AppID))
	return NewLogicClient(cc), nil
}

// func newLogicClient(c *conf.RPCClient) logic.LogicClient {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Dial))
// 	defer cancel()
// 	conn, err := grpc.DialContext(ctx, "discovery://default/goim.logic",
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
// 			grpc.WithBalancerName(roundrobin.Name),
// 		}...)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return logic.NewLogicClient(conn)
// }

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
