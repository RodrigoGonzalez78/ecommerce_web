package middleware

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func CheckJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			http.Error(w, "No autorizado - No se encontró el token", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		// Procesar el token
		_, valid, _, err := utils.ProcessToken(tokenString)
		if err != nil || !valid {
			http.Error(w, "No autorizado - Token inválido", http.StatusUnauthorized)
			return
		}

		// Pasar la solicitud al siguiente manejador
		next.ServeHTTP(w, r)
	}
}
