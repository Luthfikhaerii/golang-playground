package http

import (
	"golang-playground/internal/delivery/http/middlewares"
	"golang-playground/internal/delivery/http/routes"
	"golang-playground/internal/messaging"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SettupRoute(db *gorm.DB, producer *messaging.KafkaProducer) *gin.Engine {
	//init route
	r := gin.Default()

	//global
	r.Use(middlewares.LoggerMiddleware())
	r.Static("/upload", "./upload")

	//route
	routes.TestRoute(r)
	api := r.Group("/api")
	{
		routes.UploadRoute(api)
		routes.UserRoute(api, db)
		routes.OrderRouter(api, db, producer)
	}
	return r
}
