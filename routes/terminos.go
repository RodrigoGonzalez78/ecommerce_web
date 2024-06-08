package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func Terminos(w http.ResponseWriter, r *http.Request) {

	utils.RenderTemplate(w, "templates/front/termsanduses.html", nil)
}
