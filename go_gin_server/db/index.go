package db

import (
	"context"

	"fmt"
	"os"

	"github.com/BNsrujan/go_Server_Set.git/config"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDB() {
	DATABASE_URL := config.Config.DB_PATH
	var err error
	DB, err = pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Data bass is not connected %v\n", err)
		os.Exit(1)
	}

	errs := DB.Ping(context.Background())
	if errs != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed %v\n", err)
		os.Exit(1)
	}

	fmt.Print("file running")
}
