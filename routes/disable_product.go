package routes

import (
	"net/http"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
)

func DisableProduct(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"down": "SI",
	}

	err = db.UpdateProduct(uint(id), data)

	if err != nil {

		http.Error(w, "Unable to update product", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
