package main

import (
	"log"

	"github.com/gokuls-codes/go-echo-starter/internal/database"
	"github.com/gokuls-codes/go-echo-starter/internal/server"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(":8080", db)
	err = server.Start()
	
	if err != nil {
		log.Fatal(err)
	}
}