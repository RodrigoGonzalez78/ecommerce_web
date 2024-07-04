package routes

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

// Handler para incrementar la cantidad de un producto en el carrito
func IncrementCartItem(w http.ResponseWriter, r *http.Request) {
	modifyCartItem(w, r, 1)
}

// Handler para decrementar la cantidad de un producto en el carrito
func DecrementCartItem(w http.ResponseWriter, r *http.Request) {
	modifyCartItem(w, r, -1)
}

// Handler para eliminar un producto del carrito
func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Leer la cookie del carrito
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

	// Eliminar el producto del carrito
	newCart := []models.CartItem{}
	for _, item := range cart {
		if item.ID != id {
			newCart = append(newCart, item)
		}
	}

	// Serializar el carrito a JSON
	cartBytes, err = json.Marshal(newCart)
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

	http.Redirect(w, r, "/my-cart", http.StatusSeeOther)
}

// Función para modificar la cantidad de un producto en el carrito
func modifyCartItem(w http.ResponseWriter, r *http.Request, qtyChange int) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Leer la cookie del carrito
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

	// Buscar el producto en el carrito
	itemIndex := -1
	for i, item := range cart {
		if item.ID == id {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 {
		http.Error(w, "Producto no encontrado en el carrito", http.StatusNotFound)
		return
	}

	// Modificar la cantidad del producto
	cart[itemIndex].Qty += qtyChange
	if cart[itemIndex].Qty <= 0 {
		// Si la cantidad es 0 o menor, eliminar el producto del carrito
		cart = append(cart[:itemIndex], cart[itemIndex+1:]...)
	}

	// Serializar el carrito a JSON
	cartBytes, err = json.Marshal(cart)
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

	http.Redirect(w, r, "/my-cart", http.StatusSeeOther)
}
