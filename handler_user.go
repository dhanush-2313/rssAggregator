package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dhanush-2313/rssAggregator/internal/auth"
	"github.com/dhanush-2313/rssAggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, "Error parasing JSON")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, 500, "Error creating user")
		return
	}

	RespondWithJSON(w, 201, DatabaseUsertoUser(user))
}

func (apiCfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
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
	RespondWithJSON(w, 200, DatabaseUsertoUser(user))

}
