package main

import (
	"net/http"

	"github.com/dhanush-2313/rssAggregator/internal/auth"
	"github.com/dhanush-2313/rssAggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			RespondWithError(w, 403, "Unauthorized")
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			RespondWithError(w, 500, "Error getting user")
			return
		}
		handler(w, r, user)
	}
}
