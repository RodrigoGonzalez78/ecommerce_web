package routes

import (
	"net/http"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
)

func AttendedConsult(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"attended": "SI",
	}

	if err := db.UpdateConsult(uint(id), data); err != nil {
		http.Error(w, "Error al atender el  la consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/consults-list", http.StatusSeeOther)
}
