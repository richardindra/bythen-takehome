package blog

import (
	"bythen-takehome/pkg/response"
	"errors"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) DecodeJWT(w http.ResponseWriter, r *http.Request) {
	var (
		token []string
		err   error
		resp  response.Response

		ctx  = r.Context()
		ctxD = make(map[string]interface{})
	)
	defer resp.RenderJSON(w, r)

	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		err = errors.New("authorization header is required")
		resp.SetError(err, http.StatusUnauthorized)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)

		resp.StatusCode = http.StatusUnauthorized
		return
	}

	token = strings.Split(authorization, " ")
	if token[0] != "Bearer" {
		err = errors.New("invalid token type")
		resp.SetError(err, http.StatusUnauthorized)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)

		resp.StatusCode = http.StatusUnauthorized
		return
	}

	res, err := h.authSvc.DecodeJWT(ctx, token[1])
	if err != nil {
		resp.SetError(err, http.StatusOK)

		ctxD["error"] = err.Error()
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	resp.Data = res

	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
