package routes

import (
	"net/http"
	"time"
)

// Logout maneja el cierre de sesión de usuarios
func Logout(w http.ResponseWriter, r *http.Request) {
	// Crear una cookie con una fecha de expiración en el pasado para eliminarla
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	}

	// Establecer la cookie en el encabezado de respuesta
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
