package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

// CheckJwt es un middleware para verificar y procesar tokens JWT
func CheckJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Intentar obtener la cookie "token" del request
		cookie, err := r.Cookie("token")

		if err != nil {
			// Si no se puede obtener la cookie, crear un contexto con datos vacíos
			voidClaim := &models.Claim{
				RolID:    0,
				ID:       0,
				Email:    "",
				Name:     "",
				LastName: "",
			}
			ctx := context.WithValue(r.Context(), "userData", voidClaim)
			fmt.Println(1)
			// Pasar la solicitud al siguiente manejador con el contexto actualizado
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// Obtener el valor de la cookie
		tokenString := cookie.Value

		// Procesar el token JWT
		claim, valid, _, err := utils.ProcessToken(tokenString)

		if err != nil || !valid {
			// Si hay un error al procesar el token o no es válido, usar datos vacíos
			voidClaim := &models.Claim{
				RolID:    0,
				ID:       0,
				Email:    "",
				Name:     "",
				LastName: "",
			}
			ctx := context.WithValue(r.Context(), "userData", voidClaim)
			fmt.Println(2)
			// Pasar la solicitud al siguiente manejador con el contexto actualizado
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// Crear un nuevo contexto con los datos del usuario obtenidos del token
		ctx := context.WithValue(r.Context(), "userData", claim)
		fmt.Println(3)
		fmt.Println(claim)
		// Pasar la solicitud al siguiente manejador con el contexto actualizado
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
