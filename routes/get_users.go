package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // Valor predeterminado
	}

	users, totalUsers, err := db.GetPaginatedUsers(page, pageSize)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Calcular el número total de páginas
	totalPages := (totalUsers + int64(pageSize) - 1) / int64(pageSize)

	// Respuesta JSON con usuarios y datos de paginación
	response := map[string]interface{}{
		"users":       users,
		"totalUsers":  totalUsers,
		"totalPages":  totalPages,
		"currentPage": page,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
