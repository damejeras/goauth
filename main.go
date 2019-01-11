package main

import (
	"log"
	"net/http"

	"github.com/damejeras/goauth/api"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := api.Config{
		KeyPath: "keys/jwtRS256.key",
	}
	server := api.NewServer(config)
	log.Fatal(http.ListenAndServe(":8080", server))
}
