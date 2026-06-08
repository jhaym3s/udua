package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jhaym3s/udua/cmd"
	"github.com/jhaym3s/udua/internal/env"
)


func main(){
	fmt.Println("Hello world")

	config := cmd.Config{
		Addr: ":8080",
		DBConfig: cmd.DBConfig{
			Dsn: env.GetEnv("DB_DSN", "host=localhost user=postgres password=postgres dbname=udua sslmode=disable"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger) // this is the standard library logger, so you can use slog.Info, slog.Error, etc. anywhere in your code


	//Database 
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.DBConfig.Dsn)
	if err != nil {
		log.Fatal("database connection failed:", err)
	}
	defer conn.Close(ctx) 

	logger.Info("Connected to database", "dsn", config.DBConfig.Dsn)

	api := cmd.Application{
		Config: config,
	}

	h := api.Mount()

	
	if err := api.Run(h); err != nil {
		slog.Error("Server failed to start", err)
		log.Fatal("Server failed to start:", err)
	}




}