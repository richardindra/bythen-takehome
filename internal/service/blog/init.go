package blog

import (
	"bythen-takehome/internal/entity/blog"
	"context"
)

type Data interface {
	// User
	CreateUser(ctx context.Context, user blog.User) (int64, error)
	CheckUser(ctx context.Context, username, email string) (int, error)
	GetUserByUsername(ctx context.Context, username string) (blog.User, error)
	UpdateLastLogin(ctx context.Context, username string) error
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
