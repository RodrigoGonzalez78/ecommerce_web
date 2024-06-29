package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func ContactMePage(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("userData").(*models.Claim)

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to process request", http.StatusInternalServerError)
			return
		}

		consult := models.Consult{
			Name:        r.FormValue("name"),
			Email:       r.FormValue("email"),
			Description: r.FormValue("description"),
			Attended:    "NO",
			Archived:    "NO",
		}

		err = db.CreateConsult(consult)

		if err != nil {

			data := map[string]interface{}{
				"Titulo":    "Home",
				"IDProfile": userData.RolID,
				"Success":   false,
				"Error":     true,
			}
			utils.RenderTemplate(w, "templates/back/consults/contact.html", data)
			return
		}

		data := map[string]interface{}{
			"Titulo":    "Home",
			"IDProfile": userData.RolID,
			"Success":   true,
			"Error":     false,
		}

		utils.RenderTemplate(w, "templates/back/consults/contact.html", data)
		return
	}

	data := map[string]interface{}{
		"Titulo":    "Home",
		"IDProfile": userData.RolID,
		"Success":   false,
		"Error":     false,
	}
	utils.RenderTemplate(w, "templates/back/consults/contact.html", data)
}
