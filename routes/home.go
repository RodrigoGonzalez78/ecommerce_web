package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"Success": false,
		"Error":   false,
	}
	utils.RenderTemplate(w, "templates/front/home.html", data)
}
