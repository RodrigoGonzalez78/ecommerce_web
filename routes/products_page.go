package routes

import (
	"net/http"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func ProductsPage(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("userData").(*models.Claim)

	search := r.URL.Query().Get("search")
	pageParam := r.URL.Query().Get("page")
	page := 1

	if pageParam != "" {
		page, _ = strconv.Atoi(pageParam)
	}

	itemsPerPage := 10

	products, totalProducts, err := db.GetPaginatedProducts(search, page, itemsPerPage)
	if err != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}

	totalPages := int((totalProducts + int64(itemsPerPage) - 1) / int64(itemsPerPage))

	data := map[string]interface{}{
		"Titulo":      "Products",
		"IDProfile":   userData.RolID,
		"Products":    products,
		"Search":      search,
		"CurrentPage": page,
		"TotalPages":  totalPages,
		"Success":     false,
		"Error":       false,
		"NextPage":    page + 1,
		"Antpage":     page - 1,
	}

	utils.RenderTemplate(w, "templates/back/products/products.html", data)
}
