package handler

import (
	"github.com/LuxAeterna-git/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input jwt.User

	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
	}
	accessToken, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	refreshToken, err := h.services.Authorization.GenerateRefreshToken()
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{"accessToken": accessToken, "refreshToken": refreshToken})
}

func (h *Handler) Hello(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]interface{}{"Hello!": "Its working!"})
}
