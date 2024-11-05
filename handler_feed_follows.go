package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dhanush-2313/rssAggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
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

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
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

	RespondWithJSON(w, 201, DatabaseFeedFollowtoFeedFollow(feed_follow))
}

func (apiCfg *apiConfig) HandlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, 500, "Error fetching feed follows")
		return
	}

	RespondWithJSON(w, 201, DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		RespondWithError(w, 400, "Could not parse feed follow id")
		return
	}

	err = apiCfg.DB.Deletefeed(r.Context(), database.DeletefeedParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		RespondWithError(w, 500, "Could not delete feed follow")
		return
	}

	RespondWithJSON(w, 200, map[string]string{"message": "Feed follow deleted successfully"})
}
