package main

import (
	"log"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/routes"
)

func main() {

	// Configura el manejador de archivos estáticos para servir archivos desde el directorio "assets"
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", routes.HomeHandler)

	log.Println("Servidor iniciado en el puerto 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
