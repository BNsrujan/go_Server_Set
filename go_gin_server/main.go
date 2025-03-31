package main

import (
	"context"
	"net/http"

	"github.com/BNsrujan/go_Server_Set.git/config"
	"github.com/BNsrujan/go_Server_Set.git/db"
	"github.com/BNsrujan/go_Server_Set.git/router"
	// "github.com/gin-gonic/gin"
)

// first go will run this file until here when it import all the file the it will run init function and after that it will run the main

func main() {

	handler := router.MountRout()

	db.InitDB()
	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}
	defer db.DB.Close(context.Background())

	handler.Run()
	server.ListenAndServe()
}
