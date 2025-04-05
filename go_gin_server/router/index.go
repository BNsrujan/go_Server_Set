package router

import (
	"net/http"

	"github.com/BNsrujan/go_Server_Set.git/router/handlers"
	"github.com/gin-gonic/gin"
)

func MountRout() *gin.Engine {
	handler := gin.Default()
	// if we call this function in order that is not inlined then it may give an err so we alwase create init function
	// config.Config.LoadConfig()

	handler.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok for Gin",
		})
	})

	taskRouters := handler.Group("/task")
	{
		taskRouters.POST("/k", handlers.SaveTask)
		// taskRouters.GET("/", handlers.ReadTask)
		// taskRouters.DELETE("/", handlers.ReadTask)
	}

	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Router not Found",
		})
	})
	return handler

}
