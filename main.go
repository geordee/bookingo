package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Booking App!")
	})

	log.Println("Starting server on :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
