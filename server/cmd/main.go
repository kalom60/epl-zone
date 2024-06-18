package main

import (
	"fmt"
	"log"

	"github.com/kalom60/epl-zone/internal/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("Error initializing server: %s", err)
	}

	httpServer := server.NewHTTPServer()

	err = httpServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
