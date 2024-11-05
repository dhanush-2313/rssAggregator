package middleware

import (
	"net/http"

	config "github.com/dhanush-2313/rssAggregator/config"
	handlers "github.com/dhanush-2313/rssAggregator/handlers"
	"github.com/dhanush-2313/rssAggregator/internal/auth"
	"github.com/dhanush-2313/rssAggregator/internal/database"
)

type AuthedHandler func(apiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request, user database.User)

func MiddlewareAuth(ApiCfg *config.ApiConfig, handler AuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			handlers.RespondWithError(w, 403, "Unauthorized")
			return
		}

		user, err := ApiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			handlers.RespondWithError(w, 500, "Error getting user")
			return
		}
		handler(ApiCfg, w, r, user)
	}
}
