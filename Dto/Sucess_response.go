package Dto

import "github.com/gin-gonic/gin"

type SuccessResponseDto[T any] struct {
	Success bool   `json:"success"`
	Alert   string `json:"alert"`
	Data    *T     `json:"data,omitempty"`
}

func SuccessResponse[T any](ctx *gin.Context, alert string, data *T, status int) {
	ctx.JSON(status, SuccessResponseDto[T]{
		Success: true,
		Alert:   alert,
		Data:    data,
	})
}
