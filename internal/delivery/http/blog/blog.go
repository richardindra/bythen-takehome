package blog

import (
	"bythen-takehome/pkg/response"
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	httpHelper "bythen-takehome/internal/delivery/http"

	"github.com/gorilla/mux"
)

func (h *Handler) GetBlogByID(w http.ResponseWriter, r *http.Request) {
	var (
		resp response.Response
		ctx  = r.Context()
		ctxD = make(map[string]interface{})
	)

	defer resp.RenderJSON(w, r)

	v := mux.Vars(r)
	blogID := v["id"]
	id, _ := strconv.ParseInt(blogID, 10, 64)

	res, err := h.blogSvc.GetBlogByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, httpHelper.ErrDataNotFound):
			resp.SetError(err, http.StatusNotFound)
		default:
			resp.SetError(err, http.StatusInternalServerError)
		}

		ctxD["error"] = err.Error()
		ctx = context.WithValue(ctx, "data", ctxD)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)

		return
	}

	resp.Data = res
	resp.SetError(err, http.StatusOK)
	
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
