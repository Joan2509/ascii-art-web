package main

import (
	art "ascii/ascii"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// PageData represents data to be rendered in HTML templates
type PageData struct {
	Result string
	Error  string
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/ascii-art", handleAsciiArt)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
		}
	}
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Printf("Error parsing form: %v", err)
			return
		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")

		result, err := art.Combinator(art.Ascii(text, banner))
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			log.Printf("Error generating ASCII art: %v", err)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error parsing template: %v", err)
			return
		}

		data := PageData{
			Result: result,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Printf("Error executing template: %v", err)
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
