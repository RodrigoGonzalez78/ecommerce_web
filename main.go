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

	http.HandleFunc("/", middleware.CheckJwt(routes.HomeHandler))
	http.HandleFunc("/termsanduses", middleware.CheckJwt(routes.Terminos))
	http.HandleFunc("/about", middleware.CheckJwt(routes.About))

	http.HandleFunc("/login-page", middleware.CheckJwt(routes.LoginPage))
	http.HandleFunc("/sign-up-page", middleware.CheckJwt(routes.SignUpPage))
	http.HandleFunc("/logout", middleware.CheckJwt(routes.Logout))

	http.HandleFunc("/contact-me-page", middleware.CheckJwt(routes.ContactMePage))

	http.HandleFunc("/products-page", middleware.CheckJwt(routes.ProductsPage))
	http.HandleFunc("/new-product", middleware.CheckJwt(middleware.AdminCheck(routes.NewProduct)))
	http.HandleFunc("/edit-product", middleware.CheckJwt(middleware.AdminCheck(routes.EditProduct)))
	http.HandleFunc("/disabled-products", middleware.CheckJwt(middleware.AdminCheck(routes.DisabledProductsPage)))

	log.Println("Servidor iniciado en el puerto 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
