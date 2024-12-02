package routes

import (
	"net/http"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
)

func ArchiveConsult(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"archived": "SI",
	}

	if err := db.UpdateConsult(uint(id), data); err != nil {
		http.Error(w, "Error al archivar   la consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/consults-list", http.StatusSeeOther)
}
