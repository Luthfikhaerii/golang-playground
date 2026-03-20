package handler

import (
	"golang-playground/internal/dto"
	"golang-playground/internal/usecase"
	"golang-playground/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: *usecase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("Invalid request", zap.Error(err))

		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	accessToken, err := h.usecase.Create(req)
	if err != nil {
		logger.Log.Error("Register failed", zap.Error(err))

		c.JSON(400, gin.H{"message": "regiter failed"})
		return
	}

	c.SetCookie("accessToken", accessToken.Token, 3600*24, "/", "", false, true)

	c.JSON(200, gin.H{
		"accessToken": accessToken.Token,
	})
}
