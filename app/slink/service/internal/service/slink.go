package service

import (
	"context"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
)

func (s *ShortLinkService) CreateShortLink(ctx context.Context, req *v1.CreateShortLinkRequest) (*v1.CreateShortLinkReply, error) {
	res, err := s.su.Create(ctx, req.GetLink())
	if err != nil {
		return nil, err
	}

	return &v1.CreateShortLinkReply{
		Key: res.Key,
	}, nil
}
