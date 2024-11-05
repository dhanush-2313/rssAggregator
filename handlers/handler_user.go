package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	config "github.com/dhanush-2313/rssAggregator/config"
	models "github.com/dhanush-2313/rssAggregator/models"

	"github.com/dhanush-2313/rssAggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerCreateUser(apiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request) {
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

	RespondWithJSON(w, 201, models.DatabaseUsertoUser(user))
}

func HandlerGetUser(apiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJSON(w, 200, models.DatabaseUsertoUser(user))
}

func HandlerGetPostsForUser(apiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		RespondWithError(w, 400, "Could not get posts")
		return
	}

	RespondWithJSON(w, 200, posts)
}
