package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func ContactMePage(w http.ResponseWriter, r *http.Request) {

	utils.RenderTemplate(w, "templates/back/consults/contact.html", nil)
}
