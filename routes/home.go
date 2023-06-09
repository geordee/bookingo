package routes

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Booking App</h1>")
	fmt.Fprintf(w, "<p>Welcome to Booking App!</p>")
}
