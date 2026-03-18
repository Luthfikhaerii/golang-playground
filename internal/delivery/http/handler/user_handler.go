package handler

import (
	"golang-playground/internal/dto"
	"golang-playground/internal/usecase"

	"github.com/gin-gonic/gin"
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
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	accessToken, err := h.usecase.Create(req)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.SetCookie("accessToken", accessToken.Token, 3600*24, "/", "", false, true)

	c.JSON(200, gin.H{
		"accessToken": accessToken,
	})
}
