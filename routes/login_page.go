package routes

import (
	"net/http"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
	"github.com/golang-jwt/jwt/v5"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.RenderTemplate(w, "templates/back/users/login.html", nil)
		return
	}

	if 40 != 30 {
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString("sdsd")

		if err != nil {
			http.Error(w, "Error al generar el token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
		})

		w.Write([]byte("Inicio de sesión exitoso y cookie establecida!"))
	} else {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
	}

}
