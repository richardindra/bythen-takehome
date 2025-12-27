package blog

import (
	"bythen-takehome/internal/entity/auth"
	"bythen-takehome/internal/entity/blog"
	"context"
)

type authSvc interface {
	Register(ctx context.Context, user blog.User) (blog.RespCreateUser, error)
	Login(ctx context.Context, req blog.LoginRequest) (auth.LoginResponse, blog.UserInfo, error)

	GetJWTDetail(authToken string) (auth.DecodeJWT, error)
}

type (
	Handler struct {
		authSvc authSvc
	}
)

func New(authSvc authSvc) *Handler {
	return &Handler{
		authSvc: authSvc,
	}
}
