package handler

import (
	"net/http"

	model "github.com/alnszzx/HamburgueriaGo/Database/Model"
	service "github.com/alnszzx/HamburgueriaGo/Service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SnacksHandler struct {
	service *service.SnacksService
}

func NewSnacksHandler(service *service.SnacksService) *SnacksHandler {
	return &SnacksHandler{
		service: service,
	}
}

func (sh *SnacksHandler) CreateSnackHandler(ginContext *gin.Context) {
	var snack model.Snacks

	if err := ginContext.ShouldBindJSON(&snack); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}

	if err := sh.service.CreateSnackService(&snack); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"alert": "error to create a new Snack",
			"error": err.Error(),
		})
		return
	}

	ginContext.JSON(http.StatusCreated, snack)
}

func (sh *SnacksHandler) GetSnackByIdHandler(ginContext *gin.Context) {
	idParam := ginContext.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}

	snack, err := sh.service.GetSnackByIdService(id)
	if err != nil {
		ginContext.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginContext.JSON(http.StatusOK, snack)
}

func (sh *SnacksHandler) DeleteSnackHandler(ginContext *gin.Context) {
	idParam := ginContext.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}

	if err := sh.service.DeleteSnackService(id); err != nil {
		ginContext.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginContext.JSON(http.StatusNoContent, "delete employer sucessfuly")
}

func (sh *SnacksHandler) UpdateSnackHandler(ginContext *gin.Context) {
	idParam := ginContext.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}

	var snack model.Snacks

	if err := ginContext.ShouldBindJSON(&snack); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := sh.service.UpdateSnackService(id, &snack); err != nil {
		ginContext.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Snack updated successfully",
	})
}
