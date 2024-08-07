package utils

import (
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {

	// Cargar los archivos HTML para el encabezado, navegación y pie de página
	headerTemplate, err := template.ParseFiles("templates/front/head_view.html")

	if err != nil {
		http.Error(w, "No se pudo cargar el archivo de encabezado", http.StatusInternalServerError)
		return
	}

	navTemplate, err := template.ParseFiles("templates/front/nav_view.html")

	if err != nil {
		http.Error(w, "No se pudo cargar el archivo de navegación", http.StatusInternalServerError)
		return
	}

	contentTemplate, err := template.ParseFiles(tmpl)

	if err != nil {
		http.Error(w, "No se pudo cargar el archivo de contenido"+err.Error(), http.StatusInternalServerError)
		return
	}

	footerTemplate, err := template.ParseFiles("templates/front/footer_view.html")
	if err != nil {
		http.Error(w, "No se pudo cargar el archivo de pie de página", http.StatusInternalServerError)
		return
	}

	// Renderizar los archivos HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := headerTemplate.Execute(w, data); err != nil {
		http.Error(w, "No se pudo renderizar el encabezado", http.StatusInternalServerError)
		return
	}

	if err := navTemplate.Execute(w, data); err != nil {
		http.Error(w, "No se pudo renderizar la navegación", http.StatusInternalServerError)
		return
	}

	if err := contentTemplate.Execute(w, data); err != nil {
		http.Error(w, "No se pudo renderizar el contenido", http.StatusInternalServerError)
		return
	}

	if err := footerTemplate.Execute(w, nil); err != nil {
		http.Error(w, "No se pudo renderizar el pie de página", http.StatusInternalServerError)
		return
	}

}

func FileOnlyHandler(root http.Dir) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(string(root), r.URL.Path)
		info, err := os.Stat(path)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// Verificar si la ruta es un archivo
		if info.IsDir() {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, path)
	})
}

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// HashPassword encripta la contraseña usando bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compara una contraseña en texto claro con su hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
