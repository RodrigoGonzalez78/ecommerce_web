package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func UserList(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)
	users, _ := db.GetAllUsers()

	data := map[string]interface{}{
		"Titulo":    "Lista de Usuarios",
		"IDProfile": userData.RolID,
		"Users":     users,
		"Success":   false,
		"Error":     false,
	}
	utils.RenderTemplate(w, "templates/back/consults/consult_list.html", data)
}
