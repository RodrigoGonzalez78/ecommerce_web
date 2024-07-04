package routes

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

// / Handler para manejar el carrito de compras
func AddToCart(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del producto de la URL
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Leer la cookie del carrito
	cartCookie, err := r.Cookie("cart")
	var cart []models.CartItem

	if err == nil {
		// La cookie existe, decodificar el contenido
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

	// Buscar el producto en el carrito
	itemIndex := -1
	for i, item := range cart {
		if item.ID == id {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 {
		// El producto no está en el carrito, añadirlo con cantidad 1
		cart = append(cart, models.CartItem{ID: id, Qty: 1})
	} else {
		// El producto ya está en el carrito, incrementar la cantidad
		cart[itemIndex].Qty++
	}

	// Serializar el carrito a JSON
	cartBytes, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, "Error al serializar el carrito", http.StatusInternalServerError)
		return
	}

	// Codificar el carrito en base64
	cartValue := base64.URLEncoding.EncodeToString(cartBytes)

	// Crear la cookie con duración de 24 horas
	expiration := time.Now().Add(24 * time.Hour)

	cartCookie = &http.Cookie{
		Name:     "cart",
		Value:    cartValue,
		Expires:  expiration,
		HttpOnly: true,
	}

	// Añadir la cookie a la respuesta
	http.SetCookie(w, cartCookie)
}
