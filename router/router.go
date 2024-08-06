package router

import (
	"github.com/AnnonaOrg/sendtome/handler/api_handler"

	"github.com/gin-gonic/gin"
	"github.com/AnnonaOrg/sendtome/handler/webhook_handler/tele_handler"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())

	g.Use(mw...)
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	g.MaxMultipartMemory = 8 << 20 // 8 MiB

	g.NoRoute(api_handler.ApiNotFound)
	g.GET("/", api_handler.ApiHello)
	g.GET("/ping", api_handler.ApiPing)

	webhookR := g.Group("/webhook")
	{
		webhookR.POST("/tele/:botToken", tele_handler.Update)
	}

	return g
}
