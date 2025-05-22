package routes

import (
	"net/http"

	"github.com/Humiditii/movie-api/internal/auth"
	"github.com/Humiditii/movie-api/internal/movies"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(r *gin.Engine, movieHandler *movies.Handler, authHandler *auth.Handler) {
	api := r.Group("/api")

	{
		api.GET("/",defaultHandler)
	}

	{
		movieRoutes := api.Group("/movies")
		movieRoutes.GET("/", movieHandler.GetMovies)
	}

	{
		authRoutes := api.Group("/auth")
		authRoutes.POST("/signup", authHandler.SignupUser)
	}

}


func defaultHandler( c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"message":"hey man",
	})
}
