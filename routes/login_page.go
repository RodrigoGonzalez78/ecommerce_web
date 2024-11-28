package routes

import (
	"net/http"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

const (
	errInvalidCredentials = "Credenciales inválidas"
	errGeneratingToken    = "Error al generar el token"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)

	data := map[string]interface{}{
		"Titulo":    "Inicia Sesión",
		"IDProfile": userData.RolID,
		"Error":     "", // Añadimos un campo para el mensaje de error
	}

	// Redirigir si el usuario ya está autenticado
	if userData.RolID != 0 {
		http.Redirect(w, r, "/home-page", http.StatusSeeOther)
		return
	}

	// Si el método no es POST, se renderiza la página de login
	if r.Method != http.MethodPost {
		utils.RenderTemplate(w, "templates/back/users/login.html", data)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validación de entrada
	if email == "" || password == "" {
		data["Error"] = errInvalidCredentials // Agregar el error a los datos
		utils.RenderTemplate(w, "templates/back/users/login.html", data)
		return
	}

	// Verificar si el usuario existe
	exists, user, err := db.CheckExistUser(email)
	if err != nil {
		data["Error"] = "Error en el servidor" // Error del servidor
		utils.RenderTemplate(w, "templates/back/users/login.html", data)
		return
	}

	if !exists || !utils.CheckPasswordHash(password, user.Password) {
		data["Error"] = errInvalidCredentials // Error de credenciales inválidas
		utils.RenderTemplate(w, "templates/back/users/login.html", data)
		return
	}

	// Generar token
	tokenString, err := utils.CreateToken(user)
	if err != nil {
		data["Error"] = errGeneratingToken // Error generando el token
		utils.RenderTemplate(w, "templates/back/users/login.html", data)
		return
	}

	// Establecer cookie
	expiration := time.Now().Add(365 * 24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expiration,
		HttpOnly: true,
	})

	// Redirigir al home page
	http.Redirect(w, r, "/home-page", http.StatusSeeOther)
}
