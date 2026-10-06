package routes

import (
	handler "github.com/alnszzx/HamburgueriaGo/Handler"
	"github.com/gin-gonic/gin"
)

func EmployerRoutes(router *gin.Engine, employerHandler *handler.EmployerHandler) {
	router.Group("/employer")
	{
		router.POST("/", employerHandler.CreateEmployerHandler)
		router.DELETE("/:id", employerHandler.DeleteEmployerHandler)
	}
}
