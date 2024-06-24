package middleware

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func AdminCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userData, _ := r.Context().Value("userData").(*models.Claim)

		if userData.RolID != 1 {
			http.Redirect(w, r, "/home-page", http.StatusSeeOther)
		}
		// Pasar la solicitud al siguiente manejador
		next.ServeHTTP(w, r)
	}
}
