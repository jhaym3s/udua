package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/jhaym3s/udua/cmd"
)


func main(){
	fmt.Println("Hello world")

	config := cmd.Config{
		Addr: ":8080",
		DBConfig: cmd.DBConfig{},
	}

	api := cmd.Application{
		Config: config,
	}

	h := api.Mount()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)// this is the standard library logger, so you can use slog.Info, slog.Error, etc. anywhere in your code

	if err := api.Run(h); err != nil {
		slog.Error("Server failed to start: ", err)
		log.Println("Server Failed to start:", err)
		log.Fatal()
		os.Exit(1)
	}




}