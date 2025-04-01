package router

import (
	ping "github.com/BNsrujan/go_Server_Set.git/routers/handlers"
	"github.com/gin-gonic/gin"
)

func HandleRoutes() *gin.Engine {
	handler := gin.Default()

	
	handler.GET("/ping", ping.Ping)

	return handler
}
