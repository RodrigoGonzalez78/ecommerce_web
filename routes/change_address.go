package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func ChangeAddress(w http.ResponseWriter, r *http.Request) {
	// Obtiene los datos del usuario desde el contexto
	userData, _ := r.Context().Value("userData").(*models.Claim)

	// Obtiene los valores del formulario
	street := r.FormValue("street")
	postalCode := r.FormValue("postal_code")
	neighborhood := r.FormValue("neighborhood")
	city := r.FormValue("city")

	// Obtiene el usuario desde la base de datos
	user, err := db.GetUser(userData.ID)

	if err != nil {
		http.Error(w, "Error al obtener el usuario", http.StatusInternalServerError)
		return
	}

	// Si el usuario ya tiene una dirección, se actualiza
	if user.IDAddress != nil {

		err = db.UpdateAddress(*user.IDAddress, map[string]interface{}{
			"street":       street,
			"postal_code":  postalCode,
			"neighborhood": neighborhood,
			"city":         city,
		})

		if err != nil {
			http.Error(w, "Error al actualizar la dirección", http.StatusInternalServerError)
			return
		}

	} else {

		// Si el usuario no tiene una dirección, se crea una nueva
		address := &models.Address{
			Street:       street,
			PostalCode:   postalCode,
			Neighborhood: neighborhood,
			City:         city,
		}

		addressID, err := db.CreateAddress(address)

		if err != nil {
			http.Error(w, "Error al crear la dirección", http.StatusInternalServerError)
			return
		}

		// Actualiza el usuario para asociar la nueva dirección
		err = db.UpdateUser(user.ID, map[string]interface{}{
			"id_address": addressID,
		})
		if err != nil {
			http.Error(w, "Error al asociar la nueva dirección al usuario", http.StatusInternalServerError)
			return
		}
	}

	// Redirige al perfil del usuario
	http.Redirect(w, r, "/user-profile", http.StatusSeeOther)
}
