package blog

import (
	"bythen-takehome/internal/entity/blog"
	"bythen-takehome/pkg/response"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  = r.Context()
		ctxD = make(map[string]interface{})
	)

	defer resp.RenderJSON(w, r)

	user := blog.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp.SetError(err, http.StatusOK)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	res, err := h.authSvc.CreateUser(ctx, user)
	if err != nil {
		resp.SetError(err, http.StatusOK)
		ctxD["error"] = err.Error()
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	resp.Data = res
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
