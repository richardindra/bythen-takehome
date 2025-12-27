package blog

import (
	"bythen-takehome/internal/entity/auth"
	"bythen-takehome/internal/entity/blog"
	"context"
)

type BlogSvc interface {
	// Blog
	GetBlogByID(ctx context.Context, id int64) (blog.Blog, error)

	GetJWTDetail(authToken string) (auth.DecodeJWT, error)
}

type (
	Handler struct {
		blogSvc BlogSvc
	}
)

func New(bs BlogSvc) *Handler {
	return &Handler{
		blogSvc: bs,
	}
}
