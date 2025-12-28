package blog

import (
	"bythen-takehome/pkg/response"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	httpHelper "bythen-takehome/internal/delivery/http"
	"bythen-takehome/internal/entity/blog"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  = r.Context()
	)

	token, err := extractBearerToken(r)
	if err != nil {
		handleError(w, r, err, http.StatusUnauthorized)
		return
	}

	req := blog.Blog{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)

		return
	}

	res, err := h.blogSvc.CreateBlog(ctx, req, token)
	switch {
	case errors.Is(err, httpHelper.ErrTokenExpired):
		handleError(w, r, err, http.StatusUnauthorized)
	case err != nil:
		handleError(w, r, err, http.StatusInternalServerError)
	default:
		defer resp.RenderJSON(w, r)
		resp.Data = res
		resp.SetError(nil, http.StatusOK)
		log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	}
}

func (h *Handler) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  = r.Context()
	)

	token, err := extractBearerToken(r)
	if err != nil {
		handleError(w, r, err, http.StatusUnauthorized)
		return
	}

	v := mux.Vars(r)
	blogID := v["id"]
	id, _ := strconv.ParseInt(blogID, 10, 64)

	res, err := h.blogSvc.GetBlogByID(ctx, id, token)
	switch {
	case errors.Is(err, httpHelper.ErrDataNotFound):
		handleError(w, r, err, http.StatusNotFound)
	case errors.Is(err, httpHelper.ErrTokenExpired):
		handleError(w, r, err, http.StatusUnauthorized)
	case err != nil:
		handleError(w, r, err, http.StatusInternalServerError)
	default:
		defer resp.RenderJSON(w, r)
		resp.Data = res
		resp.SetError(nil, http.StatusOK)
		log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	}
}

func (h *Handler) GetAllBlog(w http.ResponseWriter, r *http.Request) {
	var (
		res  interface{}
		meta interface{}
		resp response.Response
		ctx  = r.Context()
	)

	token, err := extractBearerToken(r)
	if err != nil {
		handleError(w, r, err, http.StatusUnauthorized)
		return
	}

	searchBy := r.FormValue("search")
	author, _ := strconv.ParseInt(r.FormValue("author"), 10, 64)
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	if limit < 1 {
		limit = 15 // set default
	}
	page, _ := strconv.Atoi(r.FormValue("page"))
	if page < 1 {
		page = 1 // set default
	}

	switch {
	case searchBy == "":
		res, meta, err = h.blogSvc.GetAllBlog(ctx, r.FormValue("sort"), page, limit, token)
	case searchBy == "author":
		res, meta, err = h.blogSvc.GetAllBlogByAuthor(ctx, author, r.FormValue("sort"), page, limit, token)
	}

	switch {
	case errors.Is(err, httpHelper.ErrDataNotFound):
		handleError(w, r, err, http.StatusNotFound)
	case errors.Is(err, httpHelper.ErrTokenExpired):
		handleError(w, r, err, http.StatusUnauthorized)
	case err != nil:
		handleError(w, r, err, http.StatusInternalServerError)
	default:
		defer resp.RenderJSON(w, r)
		resp.Data = res
		resp.Metadata = meta
		resp.SetError(nil, http.StatusOK)
		log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	}
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  = r.Context()
	)

	token, err := extractBearerToken(r)
	if err != nil {
		handleError(w, r, err, http.StatusUnauthorized)
		return
	}

	v := mux.Vars(r)
	blogID := v["id"]
	id, _ := strconv.ParseInt(blogID, 10, 64)

	req := blog.Blog{}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)

		return
	}

	req.ID = id
	res, err := h.blogSvc.UpdatePost(ctx, req, token)
	switch {
	case errors.Is(err, httpHelper.ErrTokenExpired):
		handleError(w, r, err, http.StatusUnauthorized)
	case errors.Is(err, httpHelper.ErrUnauthorized):
		handleError(w, r, err, http.StatusUnauthorized)
	case err != nil:
		handleError(w, r, err, http.StatusInternalServerError)
	default:
		defer resp.RenderJSON(w, r)
		resp.Data = res
		resp.SetError(nil, http.StatusOK)
		log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	}
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  = r.Context()
	)

	token, err := extractBearerToken(r)
	if err != nil {
		handleError(w, r, err, http.StatusUnauthorized)
		return
	}

	v := mux.Vars(r)
	blogID := v["id"]
	id, _ := strconv.ParseInt(blogID, 10, 64)

	err = h.blogSvc.DeletePost(ctx, id, token)
	switch {
	case errors.Is(err, httpHelper.ErrTokenExpired):
		handleError(w, r, err, http.StatusUnauthorized)
	case errors.Is(err, httpHelper.ErrUnauthorized):
		handleError(w, r, err, http.StatusUnauthorized)
	case err != nil:
		handleError(w, r, err, http.StatusInternalServerError)
	default:
		defer resp.RenderJSON(w, r)
		resp.SetError(nil, http.StatusOK)
		log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	}
}
