package routes

import (
	"fmt"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)

	fmt.Println(userData)
	data := map[string]interface{}{
		"Titulo":    "Home",
		"IDProfile": userData.RolID,
		"Success":   false,
		"Error":     false,
	}

	utils.RenderTemplate(w, "templates/front/home.html", data)
}
