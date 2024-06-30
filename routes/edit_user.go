package routes

import (
	"net/http"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func EditUser(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	userData, _ := r.Context().Value("userData").(*models.Claim)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {

		user, err := db.GetUser(uint(id))

		if err != nil {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}

		data := map[string]interface{}{
			"Titulo":        "Editar Usuario",
			"IDProfile":     userData.RolID,
			"UserName":      user.Name,
			"ID":            user.ID,
			"UserLastName":  user.LastName,
			"UserEmail":     user.Email,
			"UserIDProfile": user.IDProfile,
		}
		utils.RenderTemplate(w, "templates/back/users/edit_user.html", data)

	} else if r.Method == http.MethodPost {

		name := r.FormValue("name")
		lastName := r.FormValue("last_name")
		email := r.FormValue("email")
		idProfile, err := strconv.Atoi(r.FormValue("id_profile"))

		if err != nil {
			http.Error(w, "ID de perfil inválido", http.StatusBadRequest)
			return
		}

		data := map[string]interface{}{
			"name":       name,
			"last_name":  lastName,
			"email":      email,
			"id_profile": uint(idProfile),
		}

		// Actualizar el usuario en la base de datos
		if err := db.UpdateUser(uint(id), data); err != nil {
			http.Error(w, "Error al actualizar el usuario: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/user-list", http.StatusSeeOther)

	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
