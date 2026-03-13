package routes

import "github.com/gin-gonic/gin"

func ProductRoute(r *gin.Engine) {
	productGroup := r.Group("/product")
	{
		productGroup.GET("/")
	}
}
