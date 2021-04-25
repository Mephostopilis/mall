package sso

import (
	"context"
	"time"

	grpctransport "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
)

// AppID .
const AppID = "discovery:///sso"
const target = "127.0.0.1:9002"

// NewClient new grpc client
func NewAdmin(ctx context.Context, opts ...grpctransport.ClientOption) (AdminClient, error) {
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
	return NewAdminClient(cc), nil
}

// NewClient new grpc client
func NewApi(ctx context.Context, opts ...grpctransport.ClientOption) (ApiClient, error) {
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
	return NewApiClient(cc), nil
}

// NewClient new grpc client
func NewSso(ctx context.Context, opts ...grpctransport.ClientOption) (SsoClient, error) {
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
	return NewSsoClient(cc), nil
}

// NewGrpcConn new grpc client
func NewGrpcConn(ctx context.Context, opts ...grpctransport.ClientOption) (*grpc.ClientConn, error) {
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
	return cc, nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
