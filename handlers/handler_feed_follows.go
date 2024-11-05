package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	config "github.com/dhanush-2313/rssAggregator/config"
	"github.com/dhanush-2313/rssAggregator/internal/database"
	models "github.com/dhanush-2313/rssAggregator/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func HandlerCreateFeedFollow(ApiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, "Error parasing JSON")
		return
	}

	feed_follow, err := ApiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		RespondWithError(w, 500, "Error creating feed follow")
		return
	}

	RespondWithJSON(w, 201, models.DatabaseFeedFollowtoFeedFollow(feed_follow))
}

func HandlerGetFeedFollow(ApiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := ApiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, 500, "Error fetching feed follows")
		return
	}

	RespondWithJSON(w, 201, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func HandlerDeleteFeedFollow(ApiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		RespondWithError(w, 400, "Could not parse feed follow id")
		return
	}

	err = ApiCfg.DB.Deletefeed(r.Context(), database.DeletefeedParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		RespondWithError(w, 500, "Could not delete feed follow")
		return
	}

	RespondWithJSON(w, 200, map[string]string{"message": "Feed follow deleted successfully"})
}
