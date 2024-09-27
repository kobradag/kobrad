package server

import (
	"context"
	"github.com/kobradag/kobrad/cmd/kobrawallet/daemon/pb"
	"github.com/kobradag/kobrad/version"
)
func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}