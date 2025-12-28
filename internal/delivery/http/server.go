package http

import (
	"net/http"

	"bythen-takehome/pkg/grace"

	"github.com/rs/cors"
)

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type BlogHandler interface {
	//Blog
	CreateBlog(w http.ResponseWriter, r *http.Request)
	GetBlogByID(w http.ResponseWriter, r *http.Request)
	GetAllBlog(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	Auth AuthHandler
	Blog BlogHandler
}

func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
