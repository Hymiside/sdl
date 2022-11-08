package handler

import (
	"github.com/Hymiside/stubent-media-backend/pkg/service"
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

	api := router.Group("/api", h.schoolIdentity)
	{
		api.POST("/create-class", h.createClass)
		api.POST("/create-student", h.createStudent)

		api.GET("/get-classes", h.getAllClasses)
	}

	return router
}
