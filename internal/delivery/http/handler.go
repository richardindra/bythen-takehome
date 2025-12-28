package http

import (
	"errors"
	"log"
	"net/http"

	"bythen-takehome/pkg/response"

	"github.com/gorilla/mux"
)

func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.HandleFunc("", defaultHandler).Methods("GET")
	r.HandleFunc("/", defaultHandler).Methods("GET")

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("", defaultHandler).Methods("GET")
	auth.HandleFunc("/", defaultHandler).Methods("GET")

	blog := r.PathPrefix("/blog").Subrouter()
	blog.HandleFunc("", defaultHandler).Methods("GET")
	blog.HandleFunc("/", defaultHandler).Methods("GET")

	//Auth
	authV1 := auth.PathPrefix("/v1").Subrouter()

	authV1.HandleFunc("/register", s.Auth.Register).Methods("POST")
	authV1.HandleFunc("/login", s.Auth.Login).Methods("POST")

	//Blog
	blogV1 := blog.PathPrefix("/v1").Subrouter()

	blogV1.HandleFunc("/posts", s.Blog.CreateBlog).Methods("POST")
	blogV1.HandleFunc("/posts/{id}", s.Blog.GetBlogByID).Methods("GET")
	blogV1.HandleFunc("/posts", s.Blog.GetAllBlog).Methods("GET")
	blogV1.HandleFunc("/posts/{id}", s.Blog.UpdatePost).Methods("PUT")
	blogV1.HandleFunc("/posts/{id}", s.Blog.DeletePost).Methods("DELETE")

	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Example Service API"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp   *response.Response
		err    error
		errRes response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	err = errors.New("404 Not Found")

	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   404,
			Msg:    "404 Not Found",
			Status: true,
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = 404
		resp.Error = errRes
		return
	}
}
