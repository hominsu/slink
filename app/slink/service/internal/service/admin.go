package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
)

func (s *AdminService) FlushShortLink(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.su.Flush(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *AdminService) ListShortLinks(ctx context.Context, req *v1.ListShortLinkRequest) (*v1.ListShortLinkReply, error) {
	res, nextPageToken, err := s.su.List(
		ctx,
		int(req.GetPageSize()),
		req.GetPageToken(),
	)
	if err != nil {
		return nil, err
	}
	return &v1.ListShortLinkReply{
		Keys:          res,
		NextPageToken: nextPageToken,
	}, nil
}

func (s *AdminService) DeleteShortLink(ctx context.Context, req *v1.DeleteShortLinkRequest) (*emptypb.Empty, error) {
	err := s.su.Del(ctx, req.GetKey())
	if err != nil {
		return nil, err
	}
	return nil, nil
}
