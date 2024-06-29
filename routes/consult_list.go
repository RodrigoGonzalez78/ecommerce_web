package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func ConsultList(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)
	consults, _ := db.GetAllNewConsults()

	data := map[string]interface{}{
		"Titulo":    "Lista de consultas",
		"IDProfile": userData.RolID,
		"Consults":  consults,
		"Success":   false,
		"Error":     false,
	}
	utils.RenderTemplate(w, "templates/back/consults/consult_list.html", data)
}
