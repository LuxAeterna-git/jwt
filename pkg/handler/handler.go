package handler

import (
	"github.com/LuxAeterna-git/jwt/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	test := router.Group("/test", h.userIdentity)
	{
		test.GET("/hello", h.Hello)
	}
	return router
}
