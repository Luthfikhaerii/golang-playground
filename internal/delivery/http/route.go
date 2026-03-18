package http

import (
	"fmt"
	"golang-playground/internal/database"
	"golang-playground/internal/delivery/http/routes"

	"github.com/gin-gonic/gin"
)

func SettupRoute() *gin.Engine {
	r := gin.Default()
	db := database.SettupDatabase()
	routes.TestRoute(r)
	api := r.Group("/api")
	{
		routes.UserRoute(api, db)
	}
	fmt.Println("\n=== REGISTERED ROUTES ===")
	for _, route := range r.Routes() {
		fmt.Printf("%s %s\n", route.Method, route.Path)
	}
	return r
}
