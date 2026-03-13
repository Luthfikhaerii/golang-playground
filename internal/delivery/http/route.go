package http

import (
	"golang-playground/internal/delivery/http/routes"

	"github.com/gin-gonic/gin"
)

func SettupRoute() *gin.Engine {
	r := gin.Default()

	r.Group("/api")
	{
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "hello world"})
		})
		routes.UserRoute(r)
	}
	return r
}
