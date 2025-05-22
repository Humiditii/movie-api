package movies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetMovies(c *gin.Context) {
	movies, err := h.service.ListMovies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}
	c.JSON(http.StatusOK, movies)
}
