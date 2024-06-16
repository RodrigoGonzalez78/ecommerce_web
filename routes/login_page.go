package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {

	utils.RenderTemplate(w, "templates/back/users/login.html", nil)
}
