package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oranges0da/go-server/services"
)

func Home() *gin.RouterGroup {
	r := gin.New()

	routes := r.Group("/home")
	{
		routes.GET("/", services.GetHome())
	}

	return routes
}
