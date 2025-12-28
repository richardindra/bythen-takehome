package blog

import (
	"bythen-takehome/internal/entity/auth"
	"bythen-takehome/internal/entity/blog"
	"context"

	"github.com/golang-jwt/jwt/v5"
)

type Data interface {
	// Blog
	CreateBlog(ctx context.Context, blog blog.Blog) (int64, error)
	UpdateViewCount(ctx context.Context, id int64) error
	GetBlogByID(ctx context.Context, id int64) (blog.Blog, error)
	GetAllBlog(ctx context.Context, sortType string, limit int, offset int) ([]blog.Blog, error)
	GetCountAllBlog(ctx context.Context) (int, error)
	GetAllBlogByAuthor(ctx context.Context, author int64, sortType string, limit int, offset int) ([]blog.Blog, error)
	GetCountAllBlogByAuthor(ctx context.Context, author int64) (int, error)
	UpdatePost(ctx context.Context, id int64, body blog.Blog) error
	DeletePost(ctx context.Context, id int64) error
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}

func (s Service) GetJWTDetail(authToken string) (auth.DecodeJWT, error) {
	var (
		data auth.DecodeJWT
	)

	decodeToken, _ := jwt.ParseWithClaims(authToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("tokenString"), nil
	})

	respDecode := decodeToken.Claims.(jwt.MapClaims)

	data.UserID = int64(respDecode["userid"].(float64))
	data.Username = (respDecode["username"].(string))
	data.Name = (respDecode["name"].(string))
	data.ExpireIn = int64(respDecode["exp"].(float64))

	return data, nil
}
