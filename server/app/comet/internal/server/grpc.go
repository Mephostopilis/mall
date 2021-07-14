package server

import (
	pb "edu/api/comet/grpc"
	"edu/pkg/middleware/handleerr"

	"edu/comet/internal/conf"
	"edu/comet/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger, s *service.CometService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(recovery.WithLogger(logger)),
				handleerr.Server(),
				tracing.Server(),
				logging.Server(logger),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Address != "" {
		opts = append(opts, grpc.Address(c.Grpc.Address))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	opts = append(opts, grpc.Logger(logger))
	srv := grpc.NewServer(opts...)
	pb.RegisterCometServer(srv, s)
	return srv
}
