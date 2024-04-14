package main

import (
	"net/http"

	"github.com/mkappworks/rss-aggregator-gobackend/internal/auth"
	"github.com/mkappworks/rss-aggregator-gobackend/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, "Auth error")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error getting user")
			return
		}

		handler(w, r, user)
	}
}
