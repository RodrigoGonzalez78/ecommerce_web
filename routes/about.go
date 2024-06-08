package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func About(w http.ResponseWriter, r *http.Request) {

	utils.RenderTemplate(w, "templates/front/about.html", nil)
}
