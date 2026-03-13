package handler

import (
	"golang-playground/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Register(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "hallo"})
}
