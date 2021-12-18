package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"time"

	v1 "pheidippides/api/pheidiqueue/v1"
	"pheidippides/internal/conf"
	"pheidippides/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, pheidiQueueService *service.PheidiQueueService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != 0 {
		opts = append(opts, grpc.Timeout(time.Duration(c.Grpc.Timeout*1e9)))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterPheidiQueueServer(srv, pheidiQueueService)
	return srv
}
