package handlers

import (
	"html/template"
	"log"
	"net/http"

	"asciiweb/ascii" // import a package called ascii to generate the ascii art
)

type PageData struct {
	Result string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		log.Println("Resource not found")
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Render the form
		tpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		// Create an instance of PageData with the desired result
		data := PageData{
			Result: "ASCII Art Generator", // Replace with your actual result(Ascii-art)
		}

		// Execute the template with the provided data
		if err := tpl.Execute(w, data); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:

		// Process the form and generate ASCII art
		input := r.FormValue("text")
		banner := r.FormValue("banner")

		if input == "" || banner == "" {
			log.Println("400 Bad Request: Input or Banner is empty")
			http.Error(w, "Input or Banner is empty", http.StatusBadRequest)
			return
		}

		asciiArt, err := ascii.AsciiArt(input, banner)
		if err != nil {
			log.Printf("500 Internal Server Error: Error generating ASCII art: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Result: asciiArt,
		}

		w.WriteHeader(http.StatusOK)

		log.Println("200 OK: ASCII art generated successfully")

		// Render the template with the generated ASCII art
		tpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Printf("500 Internal Server Error: Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, data)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
