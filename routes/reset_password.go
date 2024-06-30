package routes

import (
	"net/http"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func ResetPassword(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	//Reemplasar con un metodo mas apropiado
	newPassword := "12345678"

	hashedPassword, err := utils.HashPassword(newPassword)

	if err != nil {
		http.Error(w, "Error al encriptar la contraseña: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"password": hashedPassword,
	}

	if err := db.UpdateUser(uint(id), data); err != nil {
		http.Error(w, "Error al blanquear la contraseña: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user-list", http.StatusSeeOther)
}
