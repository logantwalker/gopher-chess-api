package router

import (
	"github.com/gin-gonic/gin"
	uci "github.com/logantwalker/gopher-chess-api/domain/uci_handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	route := router.Group("/")

	route.GET("/new", uci.NewGame)
	route.POST("/command", uci.Command)
	
	return router
}