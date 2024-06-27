package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func DisabledProductsPage(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)

	productList, _ := db.GetDisabledProducts()
	categoriesList, _ := db.GetCategories()

	data := map[string]interface{}{
		"Titulo":     "Products Inactivos",
		"IDProfile":  userData.RolID,
		"Products":   productList,
		"Categories": categoriesList,
		"Success":    false,
		"Error":      false,
	}

	utils.RenderTemplate(w, "templates/back/products/disabled_products.html", data)
}
