package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func SignUpPage(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)

	if userData.RolID != 0 {
		http.Redirect(w, r, "/home-page", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {

		name := r.FormValue("name")
		lastName := r.FormValue("last_name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")

		data := map[string]interface{}{
			"Titulo": "Registro",
			"values": map[string]string{
				"name":      name,
				"last_name": lastName,
				"email":     email,
				"password":  password,
			},
			"error": map[string]string{},
		}

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

		if confirmPassword == "" {
			data["error"].(map[string]string)["confirm_password"] = "Debe repetir la contraseña"
		} else if password != confirmPassword {
			data["error"].(map[string]string)["confirm_password"] = "Las contraseñas no coinciden"
		}

		if len(data["error"].(map[string]string)) > 0 {
			utils.RenderTemplate(w, "templates/back/users/sign_up.html", data)
			return
		}

		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			http.Error(w, "Error al encriptar la contraseña", http.StatusInternalServerError)
			return
		}

		err = db.CreateUser(&models.User{
			Name:      name,
			LastName:  lastName,
			Password:  hashedPassword,
			Email:     email,
			IDProfile: 2,
			Down:      "NO",
		})

		if err != nil {
			http.Error(w, "Error al registrar el usuario", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login-page", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Titulo":    "Registro",
		"IDProfile": userData.RolID,
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
