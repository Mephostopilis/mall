package uuid

import (
	"context"
	"time"

	grpctransport "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
)

// AppID .
const AppID = "service.uuid"
const target = "127.0.0.1:9004"

// NewClient new grpc client
func NewClient(ctx context.Context, opts ...grpctransport.ClientOption) (UUIDClient, error) {
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
	return NewUUIDClient(cc), nil
}

// NewGrpcConn new grpc client
func NewGrpcConn(ctx context.Context, opts ...grpctransport.ClientOption) (*grpc.ClientConn, error) {
	cc, err := grpctransport.DialInsecure(ctx,
		grpctransport.WithEndpoint(target),
		// grpctransport.WithOptions(grpc.WithInsecure()),
	)
	if err != nil {
		return nil, err
	}
	// cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	// cc, err := client.Dial(context.Background(), fmt.Sprintf("etcd://default/%s", AppID))
	return cc, nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
