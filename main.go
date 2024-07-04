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

	//Paginas estaticas
	http.HandleFunc("/", middleware.CheckJwt(routes.HomePage))
	http.HandleFunc("/termsanduses", middleware.CheckJwt(routes.Terminos))
	http.HandleFunc("/about", middleware.CheckJwt(routes.About))

	//Gestion de usuarios y autenticacion
	http.HandleFunc("/login-page", middleware.CheckJwt(routes.LoginPage))
	http.HandleFunc("/sign-up-page", middleware.CheckJwt(routes.SignUpPage))
	http.HandleFunc("/logout", middleware.CheckJwt(routes.Logout))
	http.HandleFunc("/user-list", middleware.CheckJwt(middleware.AdminCheck(routes.UserList)))
	http.HandleFunc("/desactivate-user", middleware.CheckJwt(middleware.AdminCheck(routes.DeactivateUser)))
	http.HandleFunc("/activate-user", middleware.CheckJwt(middleware.AdminCheck(routes.ActivateUser)))
	http.HandleFunc("/reset-password", middleware.CheckJwt(middleware.AdminCheck(routes.ResetPassword)))
	http.HandleFunc("/edit-user", middleware.CheckJwt(middleware.AdminCheck(routes.EditUser)))

	//Gestion de productos
	http.HandleFunc("/products-page", middleware.CheckJwt(routes.ProductsPage))
	http.HandleFunc("/new-product", middleware.CheckJwt(middleware.AdminCheck(routes.NewProduct)))
	http.HandleFunc("/edit-product", middleware.CheckJwt(middleware.AdminCheck(routes.EditProduct)))
	http.HandleFunc("/disabled-products", middleware.CheckJwt(middleware.AdminCheck(routes.DisabledProductsPage)))
	http.HandleFunc("/disable-product", middleware.CheckJwt(middleware.AdminCheck(routes.DisableProduct)))
	http.HandleFunc("/enable-product", middleware.CheckJwt(middleware.AdminCheck(routes.EnableProduct)))

	//Carrito y compra
	http.HandleFunc("/add-to-cart", middleware.CheckJwt(routes.AddToCart))
	http.HandleFunc("/my-orders", middleware.CheckJwt(routes.MyOrders))
	http.HandleFunc("/my-cart", middleware.CheckJwt(routes.MyCart))
	http.HandleFunc("/sum-cart-item", middleware.CheckJwt(routes.IncrementCartItem))
	http.HandleFunc("/rest-cart-item", middleware.CheckJwt(routes.DecrementCartItem))
	http.HandleFunc("/remove-from-cart", middleware.CheckJwt(routes.RemoveFromCart))
	http.HandleFunc("/clear-cart", middleware.CheckJwt(routes.ClearCart))

	//Gestion de contacto
	http.HandleFunc("/contact-me-page", middleware.CheckJwt(routes.ContactMePage))
	http.HandleFunc("/consults-list", middleware.CheckJwt(middleware.AdminCheck(routes.ConsultList)))
	http.HandleFunc("/archived-consults", middleware.CheckJwt(middleware.AdminCheck(routes.ArchivedConsult)))
	http.HandleFunc("/attended-consult", middleware.CheckJwt(middleware.AdminCheck(routes.AttendedConsult)))
	http.HandleFunc("/archive-consult", middleware.CheckJwt(middleware.AdminCheck(routes.ArchiveConsult)))

	log.Println("Servidor iniciado en el puerto 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
