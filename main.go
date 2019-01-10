package main

import (
	"github.com/damejeras/goauth/api"
	"log"
	"net/http"
)

func main() {
	server := api.NewServer()
	log.Fatal(http.ListenAndServe(":8080", server))
}
