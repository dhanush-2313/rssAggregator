package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	config "github.com/dhanush-2313/rssAggregator/config"
	models "github.com/dhanush-2313/rssAggregator/models"

	"github.com/dhanush-2313/rssAggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerCreateFeed(apiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
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
		UserID:    user.ID,
	})
	if err != nil {
		RespondWithError(w, 500, "Error creating feed")
		return
	}

	RespondWithJSON(w, 201, models.DatabaseFeedtoFeed(feed))
}

func HandlerGetFeed(apiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.Getfeeds(r.Context())
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintln("Error while getting feeds in HandlerGetFeed func", err))
	}

	RespondWithJSON(w, 201, models.DatabaseFeedsToStructFeeds(feeds))
}
