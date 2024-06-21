package main

import (
	"log"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/middleware"
	"github.com/RodrigoGonzalez78/ecommerce_web/routes"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func main() {

	db.DBConnection()

	assetsDir := http.Dir("assets")
	http.Handle("/assets/", http.StripPrefix("/assets/", utils.FileOnlyHandler(assetsDir)))

	http.HandleFunc("/", routes.HomeHandler)
	http.HandleFunc("/termsanduses", middleware.CheckJwt(routes.Terminos))
	http.HandleFunc("/about", routes.About)

	http.HandleFunc("/login-page", routes.LoginPage)
	http.HandleFunc("/sign-up-page", routes.SignUpPage)
	http.HandleFunc("/contact-me-page", routes.ContactMePage)

	log.Println("Servidor iniciado en el puerto 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
