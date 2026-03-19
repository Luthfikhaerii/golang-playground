package routes

import (
	"golang-playground/internal/delivery/http/handler"
	"golang-playground/internal/delivery/http/middlewares"
	"golang-playground/internal/repository"
	"golang-playground/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	handler := handler.NewUserHandler(usecase)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", middlewares.AuthMiddleware(), handler.Register)
	}
}
