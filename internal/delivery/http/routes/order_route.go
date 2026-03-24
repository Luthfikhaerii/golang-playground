package routes

import (
	"golang-playground/internal/delivery/http/handler"
	"golang-playground/internal/messaging"
	"golang-playground/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.RouterGroup, pub *messaging.KafkaPublisher) {
	repo := "repo"
	usecase := usecase.NewOrderUsecase(pub, repo)
	handler := handler.NewOrderHandler(usecase)
	userGroup := r.Group("/user")
	{
		userGroup.POST("/order", handler.Create)
	}
}
