package auth

import (
	"bythen-takehome/internal/entity/blog"
	"bythen-takehome/pkg/response"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  = r.Context()
		ctxD = make(map[string]interface{})
	)

	defer resp.RenderJSON(w, r)

	user := blog.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)

		return
	}

	res, err := h.authSvc.Register(ctx, user)
	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)

		ctxD["error"] = err.Error()
		ctx = context.WithValue(ctx, "data", ctxD)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)

		return
	}

	resp.Data = res
	resp.SetError(err, http.StatusOK)

	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  = r.Context()
		ctxD = make(map[string]interface{})
	)

	defer resp.RenderJSON(w, r)

	req := blog.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.SetError(err, http.StatusInternalServerError)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)

		return
	}

	res, meta, err := h.authSvc.Login(ctx, req)
	if err != nil {
		resp.SetError(err, http.StatusUnauthorized)

		ctxD["error"] = err.Error()
		ctx = context.WithValue(ctx, "data", ctxD)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)

		return
	}

	resp.Data = res
	resp.Metadata = meta
	resp.SetError(err, http.StatusOK)

	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
