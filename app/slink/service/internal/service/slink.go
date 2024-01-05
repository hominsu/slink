package service

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
	"github.com/hominsu/slink/app/slink/service/internal/biz"
)

func (s *ShortLinkService) CreateShortLink(ctx context.Context, req *v1.CreateShortLinkRequest) (*v1.CreateShortLinkReply, error) {
	res, err := s.su.Create(ctx, &biz.ShortLink{
		Link:     req.GetLink(),
		ExpireAt: req.GetExpireAt().AsTime(),
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateShortLinkReply{
		Key:      res.Key,
		ExpireAt: timestamppb.New(res.ExpireAt),
	}, nil
}
