package main

import (
	"log"
	"net/http"
	"os"
	"time"

	config "github.com/dhanush-2313/rssAggregator/config"
	handlers "github.com/dhanush-2313/rssAggregator/handlers"
	middleware "github.com/dhanush-2313/rssAggregator/middleware"
	utils "github.com/dhanush-2313/rssAggregator/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port not found!")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL not found!")
	}

	apiCfg := config.ConnectDB(dbUrl)

	go utils.StartScraping(
		apiCfg.DB,
		10,
		time.Minute,
	)
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandleErr)
	v1Router.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerCreateUser(apiCfg, w, r)
	})
	v1Router.Get("/users", middleware.MiddlewareAuth(apiCfg, handlers.HandlerGetUser))

	v1Router.Post("/feeds", middleware.MiddlewareAuth(apiCfg, handlers.HandlerCreateFeed))
	v1Router.Get("/feeds", middleware.MiddlewareAuth(apiCfg, handlers.HandlerGetFeed))

	v1Router.Get("/posts", middleware.MiddlewareAuth(apiCfg, handlers.HandlerGetPostsForUser))

	v1Router.Post("/feed_follows", middleware.MiddlewareAuth(apiCfg, handlers.HandlerCreateFeedFollow))
	v1Router.Get("/feed_follows", middleware.MiddlewareAuth(apiCfg, handlers.HandlerGetFeedFollow))
	v1Router.Delete("/feed_follows/{feedFollowID}", middleware.MiddlewareAuth(apiCfg, handlers.HandlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Println("Server is running on port: ", portString)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
