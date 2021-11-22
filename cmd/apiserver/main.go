package main

import (
	"log"
	"spotify-clone/server/internal/server"
)

func main() {
	var port string = "8080"
	apiserver, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(apiserver.Start(port))
}
