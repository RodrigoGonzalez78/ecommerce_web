package routes

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	cartCookie, err := r.Cookie("cart")
	if err != nil {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}

	cartBytes, err := base64.URLEncoding.DecodeString(cartCookie.Value)
	if err != nil {
		http.Error(w, "Error al decodificar el carrito", http.StatusInternalServerError)
		return
	}

	var cart []models.CartItem
	err = json.Unmarshal(cartBytes, &cart)
	if err != nil {
		http.Error(w, "Error al deserializar el carrito", http.StatusInternalServerError)
		return
	}

	newCart := []models.CartItem{}
	for _, item := range cart {
		if item.ID != id {
			newCart = append(newCart, item)
		}
	}

	cartBytes, err = json.Marshal(newCart)
	if err != nil {
		http.Error(w, "Error al serializar el carrito", http.StatusInternalServerError)
		return
	}

	cartValue := base64.URLEncoding.EncodeToString(cartBytes)

	expiration := time.Now().Add(24 * time.Hour)
	cartCookie = &http.Cookie{
		Name:     "cart",
		Value:    cartValue,
		Expires:  expiration,
		HttpOnly: true,
	}

	http.SetCookie(w, cartCookie)

	http.Redirect(w, r, "/my-cart", http.StatusSeeOther)
}
