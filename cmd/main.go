package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Humiditii/movie-api/api/routes"
	"github.com/Humiditii/movie-api/config"
	"github.com/Humiditii/movie-api/internal/auth"
	"github.com/Humiditii/movie-api/internal/movies"
	"github.com/Humiditii/movie-api/pkg/database"
	"github.com/Humiditii/movie-api/pkg/middlewares"
)

func main() {
	
	config.LoadEnv()

	db := database.Connect(config.AppEnvs.DbUrl)
	// DI for Movies
	movieRepo := movies.NewRepository(db)
	movieService := movies.NewService(movieRepo)
	movieHandler := movies.NewHandler(movieService)


	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)


	// Setup Router
	router := gin.Default()

	router.Use(middlewares.LoggingMiddleware())

	routes.RegisterRoutes(router, movieHandler, authHandler)

	log.Printf("Server running at http://localhost:%v", config.AppEnvs.Port)

	if err := http.ListenAndServe(":"+config.AppEnvs.Port, router); err != nil {
		log.Fatalf("error: %v", err)
	}
}
