package main

import (
	"fmt"
	"log"
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

	if err := api.Run(h); err != nil {
		log.Println("Server Failed to start:", err)
		log.Fatal()
		os.Exit(1)
	}




}