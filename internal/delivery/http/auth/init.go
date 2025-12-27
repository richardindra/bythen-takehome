package auth

import (
	"bythen-takehome/internal/entity/auth"
	"bythen-takehome/internal/entity/blog"
	"context"
)

type AuthSvc interface {
	Register(ctx context.Context, user blog.User) (blog.RespCreateUser, error)
	Login(ctx context.Context, req blog.LoginRequest) (auth.LoginResponse, blog.UserInfo, error)
}

type (
	Handler struct {
		authSvc AuthSvc
	}
)

func New(as AuthSvc) *Handler {
	return &Handler{
		authSvc: as,
	}
}
