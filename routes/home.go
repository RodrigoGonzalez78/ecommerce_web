package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"Success":   true,
		"Error":     false,
		"IDProfile": 1,
		"BaseURL":   "localhost:8080",
	}
	utils.RenderTemplate(w, "templates/front/home.html", data)
}
