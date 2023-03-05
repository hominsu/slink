package service

import (
	"context"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
)

func (s *ShortLinkRedirectService) Redirect(ctx context.Context, req *v1.RedirectRequest) (http.Redirector, error) {
	res, err := s.su.Get(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	return http.NewRedirect(res.Link, stdhttp.StatusMovedPermanently), nil
}
