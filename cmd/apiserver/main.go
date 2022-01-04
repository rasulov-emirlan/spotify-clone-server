package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"spotify-clone/server/config"
	"spotify-clone/server/internal/server"
	"time"
)

func main() {
	if len(os.Args) == 3 {
		config.SetSonfig(true, os.Args[2])
	}

	port, err := config.NewPortForServer()
	log.Println(port)
	if err != nil {
		log.Fatal(err)
	}
	apiserver, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := apiserver.Start(port); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := apiserver.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
