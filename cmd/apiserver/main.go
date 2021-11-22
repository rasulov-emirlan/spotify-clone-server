package main

import (
	"log"
	"spotify-clone/server/config"
	"spotify-clone/server/internal/server"
)

func main() {
	port, err := config.NewPortForServer()
	if err != nil {
		log.Fatal(err)
	}
	apiserver, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(apiserver.Start(port))
}
