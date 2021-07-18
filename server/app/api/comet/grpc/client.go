package grpc

import (
	"context"
	"time"

	grpctransport "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
)

// AppID .
const appID = "discovery:///comet"

// NewClient new grpc client
func NewClient(ctx context.Context, opts ...grpctransport.ClientOption) (CometClient, error) {
	t := make([]grpctransport.ClientOption, 0)
	t = append(t, grpctransport.WithEndpoint(appID))
	t = append(t, grpctransport.WithTimeout(time.Minute))
	t = append(t, grpctransport.WithOptions(
		grpc.KeepaliveParams(),
	))
	t = append(t, opts...)
	cc, err := grpctransport.DialInsecure(ctx, t...)
	if err != nil {
		return nil, err
	}
	return NewCometClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
