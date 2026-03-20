package http

import (
	"golang-playground/internal/database"
	"golang-playground/internal/delivery/http/middlewares"
	"golang-playground/internal/delivery/http/routes"

	"github.com/gin-gonic/gin"
)

func SettupRoute() *gin.Engine {
	r := gin.Default()
	db := database.SettupDatabase()

	//global
	r.Use(middlewares.LoggerMiddleware())
	r.Static("/upload", "./upload")

	//route
	routes.TestRoute(r)
	api := r.Group("/api")
	{
		routes.UploadRoute(api)
		routes.UserRoute(api, db)
	}
	return r
}
