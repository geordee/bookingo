package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Booking App</h1>")
		fmt.Fprintf(w, "<p>Welcome to Booking App!</p>")
	})

	log.Println("Starting server on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
