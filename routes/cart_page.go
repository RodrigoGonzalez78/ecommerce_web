package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func CartPage(w http.ResponseWriter, r *http.Request) {

	cartCookie, err := r.Cookie("cart")
	var cart []models.CartItem

	if err == nil {

		err = json.Unmarshal([]byte(cartCookie.Value), &cart)

		if err != nil {
			http.Error(w, "Error al decodificar el carrito", http.StatusInternalServerError)
			return
		}
	} else {

		cart = []models.CartItem{}
	}

}
