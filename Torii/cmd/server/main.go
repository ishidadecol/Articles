package main

import (
	"log"

	handler "github.com/Articles/Torii/internal/handler"
	"github.com/Articles/Torii/internal/server"
)

func main() {
	err := server.StartServer(":8080", handler.Echo)
	if err != nil {
		log.Fatal(err)
	}
}
