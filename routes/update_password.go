package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)

	newPassword := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if newPassword == "" || confirmPassword == "" {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	if newPassword != confirmPassword {
		http.Error(w, "Las contraseñas no coinciden", http.StatusBadRequest)
		return
	}

	if len(newPassword) < 8 {
		http.Error(w, "La contraseña debe tener al menos 8 caracteres", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		http.Error(w, "Error al encriptar la contraseña: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"password": hashedPassword,
	}

	if err := db.UpdateUser(uint(userData.ID), data); err != nil {
		http.Error(w, "Error al actualizar la contraseña: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user-profile", http.StatusSeeOther)
}
