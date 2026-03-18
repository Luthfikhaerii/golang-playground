package http

import (
	"golang-playground/internal/delivery/http/routes"

	"github.com/gin-gonic/gin"
)

func SettupRoute() *gin.Engine {
	r := gin.Default()

	routes.TestRoute(r)
	r.Group("/api")
	{
		routes.UserRoute(r)
	}
	return r
}
