package handler

import (
	"net/http"

	model "github.com/alnszzx/HamburgueriaGo/Database/Model"
	service "github.com/alnszzx/HamburgueriaGo/Service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmployerHandler struct {
	service *service.EmployerService
}

func NewEmployerHandler(service *service.EmployerService) *EmployerHandler {
	return &EmployerHandler{
		service: service,
	}
}

func (eh *EmployerHandler) CreateEmployerHandler(ginContext *gin.Context) {
	var employer model.Employer
	if err := ginContext.ShouldBindJSON(&employer); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}

	if err := eh.service.CreateEmployerService(&employer); err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{
			"alert": "Error to create a new Employer",
			"error": err.Error(),
		})
		return
	}

	ginContext.JSON(http.StatusCreated, employer)
}

func (eh *EmployerHandler) DeleteEmployerHandler(ginContext *gin.Context) {
	idParam := ginContext.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}

	if err := eh.service.DeleteEmployerService(id); err != nil {
		ginContext.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginContext.JSON(http.StatusNoContent, "delete employer sucessfuly")
}
