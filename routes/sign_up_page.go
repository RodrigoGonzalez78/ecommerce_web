package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func SignUpPage(w http.ResponseWriter, r *http.Request) {

	utils.RenderTemplate(w, "templates/back/users/sign_up.html", nil)
}
