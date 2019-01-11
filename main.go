package main

import (
	"log"
	"net/http"

	"github.com/damejeras/goauth/api"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := api.NewServer()
	log.Fatal(http.ListenAndServe(":8080", server))
}
