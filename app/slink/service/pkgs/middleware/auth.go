package middleware

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

type ContextKey string

const (

	// bearerWord the bearer key word for authorization
	bearerWord string = "Bearer"

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	authorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"

	ContextKeyUsername ContextKey = "x-md-global-username"
	ContextKeyIsAdmin  ContextKey = "x-md-global-is-admin"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(reason, "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(reason, "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(reason, "Can not get key while signing token")
)

type Claims struct {
	jwt.RegisteredClaims
	Username string
	IsAdmin  bool
}

func GenerateToken(secret, username string, isAdmin bool, expiresAt *jwt.NumericDate) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "paper4u.cn",
			Subject:   "user token",
			ExpiresAt: expiresAt,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Username: username,
		IsAdmin:  isAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JwtAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				auths := strings.SplitN(header.RequestHeader().Get(authorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
					return nil, ErrMissingJwtToken
				}
				jwtToken := auths[1]

				claims := &Claims{}
				tokenInfo, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
					if _, tOk := token.Method.(*jwt.SigningMethodHMAC); !tOk {
						return nil, ErrUnSupportSigningMethod
					}
					return []byte(secret), nil
				})
				if err != nil {
					if ve, vOk := err.(*jwt.ValidationError); vOk {
						switch {
						case ve.Errors&jwt.ValidationErrorMalformed != 0:
							return nil, ErrTokenInvalid
						case ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0:
							return nil, ErrTokenExpired
						default:
							return nil, ErrTokenParseFail
						}
					}
					return nil, errors.Unauthorized(reason, err.Error())
				} else if !tokenInfo.Valid {
					return nil, ErrTokenInvalid
				}

				// append the username and isAdmin to ctx for next service
				ctx = metadata.AppendToClientContext(ctx, string(ContextKeyUsername), claims.Username)
				ctx = metadata.AppendToClientContext(ctx, string(ContextKeyIsAdmin), strconv.FormatBool(claims.IsAdmin))
			}
			return handler(ctx, req)
		}
	}
}
