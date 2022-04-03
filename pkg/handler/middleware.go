package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithError(401, errors.New("Empty auth header"))
	}

	headerParts := strings.Split(header, " ")

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithError(401, err)
	}
	c.Set("userId", userId)

}
