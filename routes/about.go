package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func About(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("userData").(*models.Claim)

	data := map[string]interface{}{
		"Titulo":    "Home",
		"IDProfile": userData.RolID,
		"Success":   false,
		"Error":     false,
	}
	utils.RenderTemplate(w, "templates/front/about.html", data)
}
