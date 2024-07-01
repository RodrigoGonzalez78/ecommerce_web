package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

// Handler para manejar el carrito de compras
func AddToCart(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	cartCookie, err := r.Cookie("cart")
	var cart []models.CartItem

	if err == nil {

		err = json.Unmarshal([]byte(cartCookie.Value), &cart)

		if err != nil {
			http.Error(w, "Error al decodificar el carrito", http.StatusInternalServerError)
			return
		}

	}

	// Buscar el producto en el carrito
	found := false
	for i, item := range cart {
		if item.ProductID == uint(id) {
			cart[i].Quantity++
			found = true
			break
		}
	}

	// Si el producto no está en el carrito, agregarlo
	if !found {
		cart = append(cart, models.CartItem{
			ProductID: uint(id),
			Quantity:  1,
		})
	}

	cartData, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, "Error al codificar el carrito", http.StatusInternalServerError)
		return
	}

	// Crear una cookie con el carrito actualizado
	cookie := http.Cookie{
		Name:     "cart",
		Value:    string(cartData),
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	// Establecer la cookie en la respuesta
	http.SetCookie(w, &cookie)

}
