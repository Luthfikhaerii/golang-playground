package routes

import (
	"golang-playground/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	userHandler := handler.NewUserHandler()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "hello world"})
		})

		userGroup.POST("/register", userHandler.Register)
	}
}
