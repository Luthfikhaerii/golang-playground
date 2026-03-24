package routes

import (
	"golang-playground/internal/delivery/http/handler"
	"golang-playground/internal/messaging"
	"golang-playground/internal/repository"
	"golang-playground/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRouter(r *gin.RouterGroup, db *gorm.DB, pub *messaging.KafkaProducer) {
	repo := repository.NewOrderRepository(db)
	usecase := usecase.NewOrderUsecase(pub, repo)
	handler := handler.NewOrderHandler(usecase)
	r.POST("/order", handler.Create)
}
