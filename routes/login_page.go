package routes

import (
	"net/http"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

const (
	errInvalidCredentials = "Credenciales inválidas"
	errGeneratingToken    = "Error al generar el token"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RenderTemplate(w, "templates/back/users/login.html", nil)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, errInvalidCredentials, http.StatusUnauthorized)
		return
	}

	exists, user, err := db.CheckExistUser(email)

	if err != nil {
		http.Error(w, "Error en el servidor", http.StatusInternalServerError)
		return
	}

	if !exists || !utils.CheckPasswordHash(password, user.Password) {
		http.Error(w, errInvalidCredentials, http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.CreateToken(user)

	if err != nil {
		http.Error(w, errGeneratingToken, http.StatusInternalServerError)
		return
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expiration,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/home-page", http.StatusSeeOther)
}
