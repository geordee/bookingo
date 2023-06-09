package main

import (
	"log"
	"net/http"

	"github.com/geordee/bookingo/routes"
)

func main() {
	http.HandleFunc("/", routes.Home)

	log.Println("Starting server on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
