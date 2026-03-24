package handler

import (
	"golang-playground/internal/dto"
	"golang-playground/internal/usecase"
	"golang-playground/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderHandler struct {
	usecase usecase.OrderUsecase
}

func NewOrderHandler(usecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: *usecase}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("Invalid request", zap.Error(err))

		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	err := h.usecase.CreateOrder(req)
	if err != nil {
		logger.Log.Error("Register failed", zap.Error(err))

		c.JSON(400, gin.H{"message": "regiter failed"})
		return
	}

	c.JSON(200, gin.H{
		"message": "create order success",
	})
}
