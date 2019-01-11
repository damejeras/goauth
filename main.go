package main

import (
	"github.com/damejeras/goauth/api"
	"log"
	"net/http"
)

func main() {
	config := api.Config{
		KeyPath: "keys/jwtRS256.key",
	}
	server := api.NewServer(config)
	log.Fatal(http.ListenAndServe(":8080", server))
}
