package config

import (
	"database/sql"
	"log"

	"github.com/dhanush-2313/rssAggregator/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func ConnectDB(dbUrl string) *ApiConfig {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := ApiConfig{
		DB: database.New(db),
	}

	return &apiCfg
}
