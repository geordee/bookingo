package routes

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title   string
	Header  string
	Content string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	page := Page{Title: "Booking App", Header: "Booking App", Content: "Welcome to Booking App!"}
	tmpl.Execute(w, page)
}
