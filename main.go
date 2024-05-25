package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", homeHandler)

	log.Println("Servidor iniciado en el puerto 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/index.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles(tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("Error al renderizar la plantilla: %v", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
	}
}
