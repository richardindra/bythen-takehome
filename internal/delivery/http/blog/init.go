package blog

import (
	"bythen-takehome/internal/entity/auth"
	"bythen-takehome/internal/entity/blog"
	"context"
)

type authSvc interface {
	CreateUser(ctx context.Context, user blog.User) (blog.RespCreateUser, error)
	DecodeJWT(ctx context.Context, authToken string) (auth.Claims, error)
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
