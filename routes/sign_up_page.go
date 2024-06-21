package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func SignUpPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		// Recoge los datos del formulario
		name := r.FormValue("name")
		lastName := r.FormValue("last_name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Mapa para almacenar los valores y los errores
		data := map[string]interface{}{
			"values": map[string]string{
				"name":      name,
				"last_name": lastName,
				"email":     email,
				"password":  password,
			},
			"error": map[string]string{},
		}

		// Validación de los datos del formulario
		if name == "" {
			data["error"].(map[string]string)["name"] = "El nombre es obligatorio"
		}
		if lastName == "" {
			data["error"].(map[string]string)["last_name"] = "El apellido es obligatorio"
		}
		if email == "" {
			data["error"].(map[string]string)["email"] = "El correo electrónico es obligatorio"
		} else if !utils.IsValidEmail(email) {
			data["error"].(map[string]string)["email"] = "El correo electrónico no es válido"
		}

		if password == "" {
			data["error"].(map[string]string)["password"] = "La contraseña es obligatoria"
		} else if len(password) < 8 {
			data["error"].(map[string]string)["password"] = "La contraseña debe tener al menos 8 caracteres"
		}

		// Si hay errores, vuelve a mostrar el formulario con los errores
		if len(data["error"].(map[string]string)) > 0 {
			utils.RenderTemplate(w, "templates/back/users/sign_up.html", data)
			return
		}

		// Encriptar la contraseña
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			http.Error(w, "Error al encriptar la contraseña", http.StatusInternalServerError)
			return
		}

		// Aquí agregarías la lógica para registrar al usuario
		err = db.CreateUser(&models.User{
			Name:      name,
			LastName:  lastName,
			Password:  hashedPassword,
			Email:     email,
			IDProfile: 2,
		})
		if err != nil {
			http.Error(w, "Error al registrar el usuario", http.StatusInternalServerError)
			return
		}

		// Redirige al usuario a una página de éxito o inicio de sesión
		http.Redirect(w, r, "/login-page", http.StatusSeeOther)
		return
	}

	// Si el método no es POST, muestra el formulario vacío
	data := map[string]interface{}{
		"values": map[string]string{
			"name":      "",
			"last_name": "",
			"email":     "",
			"password":  "",
		},
		"error": map[string]string{},
	}

	utils.RenderTemplate(w, "templates/back/users/sign_up.html", data)
}
