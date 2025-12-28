package blog

import (
	"bythen-takehome/internal/entity/blog"
	"bythen-takehome/pkg/response"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
)

type BlogSvc interface {
	// Blog
	CreateBlog(ctx context.Context, req blog.Blog, _token string) (blog.Blog, error)
	GetBlogByID(ctx context.Context, id int64, _token string) (blog.Blog, error)
	GetAllBlog(ctx context.Context, sortType string, page int, limit int, _token string) ([]blog.Blog, interface{}, error)
	GetAllBlogByAuthor(ctx context.Context, authorID int64, sortType string, page int, limit int, _token string) ([]blog.Blog, interface{}, error)
	UpdatePost(ctx context.Context, body blog.Blog, _token string) (blog.Blog, error)
	DeletePost(ctx context.Context, id int64, _token string) error

	// Commeny
	CreateComment(ctx context.Context, req blog.Comments, _token string) (blog.Comments, error)
	GetAllCommentsByBlog(ctx context.Context, blogID int64, sortType string, page int, limit int, _token string) ([]blog.Comments, interface{}, error)
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

func extractBearerToken(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return "", errors.New("Authorization header is required")
	}

	parts := strings.Split(authorization, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("Invalid token type")
	}

	return parts[1], nil
}

func handleError(w http.ResponseWriter, r *http.Request, err error, status int) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	resp.SetError(err, status)
	log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
}
