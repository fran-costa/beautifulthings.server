package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"beautifulthings/server"
	"beautifulthings/store"
)

func main() {
	const addr = "localhost:8080"

	log.Println("Starting server at", addr)

	cancel, err := server.ServeRest(context.Background(), addr, store.NewInMemoryServer())
	if err != nil {
		log.Fatal("Error starting server: %s", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Println("Ctrl+C...")
	cancel()
}
