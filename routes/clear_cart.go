package routes

import (
	"net/http"
	"time"
)

func ClearCart(w http.ResponseWriter, r *http.Request) {
	// Crear una cookie con el mismo nombre pero con valor vacío y expiración en el pasado
	expiration := time.Now().Add(-24 * time.Hour)
	cartCookie := &http.Cookie{
		Name:     "cart",
		Value:    "",
		Expires:  expiration,
		HttpOnly: true,
	}

	// Añadir la cookie a la respuesta para eliminarla
	http.SetCookie(w, cartCookie)

	http.Redirect(w, r, "/my-cart", http.StatusSeeOther)
}
