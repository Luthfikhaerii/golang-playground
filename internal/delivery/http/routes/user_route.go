package routes

import (
	"golang-playground/internal/database"
	"golang-playground/internal/delivery/http/handler"
	"golang-playground/internal/repository"
	"golang-playground/internal/usecase"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	db := database.SettupDatabase()
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	handler := handler.NewUserHandler(usecase)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handler.Register)
	}
}
