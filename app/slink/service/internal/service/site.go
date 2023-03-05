package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
)

func (s *SiteService) Ping(_ context.Context, _ *emptypb.Empty) (*v1.PingReply, error) {
	return &v1.PingReply{
		Version: s.version,
	}, nil
}

func (s *SiteService) Count(ctx context.Context, _ *emptypb.Empty) (*v1.CountReply, error) {
	count, err := s.su.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.CountReply{Count: uint64(count)}, nil
}
