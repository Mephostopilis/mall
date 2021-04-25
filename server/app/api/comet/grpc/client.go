package grpc

import (
	"context"
	"time"

	grpctransport "github.com/go-kratos/kratos/v2/transport/grpc"
)

// AppID .
const AppID = "discovery:///comet"
const target = "direct://default/127.0.0.1:9006" // NOTE: example

// NewClient new grpc client
func NewClient(ctx context.Context, opts ...grpctransport.ClientOption) (CometClient, error) {
	t := make([]grpctransport.ClientOption, 0)
	t = append(t, grpctransport.WithEndpoint(target))
	t = append(t, grpctransport.WithTimeout(time.Minute))
	for _, o := range opts {
		t = append(t, o)
	}
	cc, err := grpctransport.DialInsecure(ctx, t...)
	if err != nil {
		return nil, err
	}
	// cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	// cc, err := client.Dial(context.Background(), fmt.Sprintf("etcd://default/%s", AppID))
	return NewCometClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
