package routes

import (
	handler "github.com/alnszzx/HamburgueriaGo/Handler"
	"github.com/gin-gonic/gin"
)

func SnackRoutes(router *gin.Engine, handler *handler.SnacksHandler) {
	snacks := router.Group("/snacks")
	{
		snacks.POST("/", handler.CreateSnackHandler)
		snacks.GET("/:id", handler.GetSnackByIdHandler)
		snacks.PUT("/:id", handler.UpdateSnackHandler)
		snacks.DELETE("/:id", handler.DeleteSnackHandler)
	}
}
