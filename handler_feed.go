package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dhanush-2313/rssAggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, "Error parasing JSON")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserId:    user.ID,
	})
	if err != nil {
		RespondWithError(w, 500, "Error creating feed")
		return
	}

	RespondWithJSON(w, 201, DatabaseFeedtoFeed(feed))
}

func (apiCfg *apiConfig) HandlerGetFeed(w http.ResponseWriter, r *http.Request, feed database.Feed) {
	RespondWithJSON(w, 200, DatabaseFeedtoFeed(feed))
}
