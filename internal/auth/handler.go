package auth

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

func (h *Handler) SignupUser(c *gin.Context) {

	data, err := h.service.SignupUser(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":err,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":"new user signup",
		"data": data,
	})


}