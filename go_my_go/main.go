package main

import (
	"context"
	"net/http"

	"github.com/BNsrujan/go_Server_Set.git/config"
	"github.com/BNsrujan/go_Server_Set.git/db"
	router "github.com/BNsrujan/go_Server_Set.git/routers"
)

func main() {

	handler := router.HandleRoutes()

	server := &http.Server{
		Addr:    config.Config.Port,
		Handler: handler,
	}
	db.DBConn()
	defer db.DB.Close(context.Background())
	// type Message struct {
	// 	Messages  string "json:messages"
	// 	CreatedAt string "json:Date"
	// }

	// handler.POST("/ping", func(ctx *gin.Context) {
	// 	var message Message
	// 	err := ctx.ShouldBindBodyWithJSON(&message)

	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "something went wrong the tasks was not added",
	// 		})
	// 		return
	// 	}

	// 	fmt.Println(message)

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "ok for Gin ", "messages": message,
	// 	})
	// })

	handler.Run()
	server.ListenAndServe()
}
