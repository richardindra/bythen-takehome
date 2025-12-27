package http

import (
	"net/http"

	"bythen-takehome/pkg/grace"

	"github.com/rs/cors"
)

type blogHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	// server  *http.Server
	Blog blogHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
