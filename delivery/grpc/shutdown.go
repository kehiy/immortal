package grpc

import (
	"context"

	rpb "github.com/dezh-tech/immortal/delivery/grpc/gen"
	"github.com/dezh-tech/immortal/pkg/logger"
)

type shutdownServer struct {
	shdCh chan struct{}
	*Server
}

func newShutdownServer(server *Server, shdCh chan struct{}) *shutdownServer {
	return &shutdownServer{
		Server: server,
		shdCh:  shdCh,
	}
}

func (s shutdownServer) Shutdown(_ context.Context, r *rpb.ShutdownRequest) (*rpb.ShutdownResponse, error) {
	logger.Info("shutdown signal received from grpc", "caller", r.String())

	s.shdCh <- struct{}{}

	return &rpb.ShutdownResponse{}, nil
}
