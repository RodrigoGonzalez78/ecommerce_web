package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func UserProfilePage(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)

	user, _ := db.GetUser(userData.ID)

	if user.IDAddress == nil {
		data := map[string]interface{}{
			"Titulo":    "Mis Datos",
			"IDProfile": userData.RolID,
			"User":      user,
			"Address":   models.Address{},
		}

		utils.RenderTemplate(w, "templates/back/users/user_profile.html", data)
		return
	}

	address, _ := db.GetAddress(*user.IDAddress)
	data := map[string]interface{}{
		"Titulo":    "Mis Datos",
		"IDProfile": userData.RolID,
		"User":      user,
		"Address":   address,
	}

	utils.RenderTemplate(w, "templates/back/users/user_profile.html", data)

}
