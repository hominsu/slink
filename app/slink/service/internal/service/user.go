package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
	"github.com/hominsu/slink/app/slink/service/pkgs/middleware"
)

func (s *UserService) Login(_ context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	if req.GetUsername() != s.secret.Admin.Username || req.GetPassword() != s.secret.Admin.Password {
		return nil, v1.ErrorUnauthorized("username or password mismatch")
	}

	expireTime := time.Now().Add(s.secret.Jwt.GetExpiredTime().AsDuration())
	token, err := middleware.GenerateToken(
		s.secret.Jwt.GetSecret(),
		req.GetUsername(),
		true,
		jwt.NewNumericDate(expireTime),
	)
	if err != nil {
		return nil, err
	}

	return &v1.LoginReply{
		Token:     token,
		ExpiresAt: timestamppb.New(expireTime),
	}, nil
}
