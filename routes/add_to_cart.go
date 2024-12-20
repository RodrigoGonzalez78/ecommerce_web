package routes

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

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

		cartBytes, err := base64.URLEncoding.DecodeString(cartCookie.Value)
		if err != nil {
			http.Error(w, "Error al decodificar el carrito", http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(cartBytes, &cart)
		if err != nil {
			http.Error(w, "Error al deserializar el carrito", http.StatusInternalServerError)
			return
		}
	}

	itemIndex := -1
	for i, item := range cart {
		if item.ID == id {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 {

		cart = append(cart, models.CartItem{ID: id, Qty: 1})
	} else {

		cart[itemIndex].Qty++
	}

	cartBytes, err := json.Marshal(cart)
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
