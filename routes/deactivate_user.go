package routes

import (
	"net/http"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
)

func DeactivateUser(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Actualizar el estado del usuario a "desactivado"
	data := map[string]interface{}{
		"down": "SI",
	}

	if err := db.UpdateUser(uint(id), data); err != nil {
		http.Error(w, "Error al desactivar el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/user-list", http.StatusSeeOther)
}
