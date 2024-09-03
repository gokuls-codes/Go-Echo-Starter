package main

import (
	"log"

	"github.com/gokuls-codes/go-echo-starter/internal/server"
)

func main() {
	server := server.NewServer(":8080", nil)
	err := server.Start()
	
	if err != nil {
		log.Fatal(err)
	}
}