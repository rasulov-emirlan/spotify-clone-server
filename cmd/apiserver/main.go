package main

import (
	"log"
	"spotify-clone/server/internal/server"
	"spotify-clone/server/internal/store"
)

func main() {
	var port string = "8080"
	s, err := store.New()
	if err != nil {
		log.Fatal(err)
	}
	apiserver, err := server.New(s)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(apiserver.Start(port))
}
