package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
	"github.com/google/uuid"
)

func NewProduct(w http.ResponseWriter, r *http.Request) {
	categoriesList, _ := db.GetCategories()

	if r.Method == http.MethodPost {
		// Collect form data
		name := r.FormValue("name")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		stock, _ := strconv.Atoi(r.FormValue("stock"))
		categoryID, _ := strconv.Atoi(r.FormValue("id_categorie"))
		description := r.FormValue("description")

		// Handle file upload
		file, handler, err := r.FormFile("image")
		if err != nil {
			fmt.Println("Error uploading file:", err)
		}
		defer file.Close()

		// Generate a random filename with the same extension
		extension := filepath.Ext(handler.Filename)
		randomFileName := uuid.New().String() + extension

		uploadDir := "assets/uploads"
		os.MkdirAll(uploadDir, os.ModePerm)
		imagePath := filepath.Join(uploadDir, randomFileName)

		dst, err := os.Create(imagePath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Map to store values and errors
		data := map[string]interface{}{
			"values": map[string]string{
				"name":         name,
				"price":        r.FormValue("price"),
				"stock":        r.FormValue("stock"),
				"id_categorie": r.FormValue("id_categorie"),
				"description":  description,
				"image":        handler.Filename,
			},
			"errors":     map[string]string{},
			"Categories": categoriesList,
		}

		// Form validation
		if name == "" {
			data["errors"].(map[string]string)["name"] = "El nombre es obligatorio"
		}
		if price <= 0 {
			data["errors"].(map[string]string)["price"] = "El precio debe ser mayor que 0"
		}
		if stock <= 0 {
			data["errors"].(map[string]string)["stock"] = "El stock debe ser mayor que 0"
		}
		if categoryID == 0 {
			data["errors"].(map[string]string)["id_categorie"] = "La categoría es requerida"
		}
		if description == "" {
			data["errors"].(map[string]string)["description"] = "La descripción es requerida"
		}
		if handler.Filename == "" {
			data["errors"].(map[string]string)["image"] = "La imagen es requerida"
		}

		// If there are errors, re-render the form with errors
		if len(data["errors"].(map[string]string)) > 0 {
			utils.RenderTemplate(w, "templates/back/products/new_product.html", data)
			return
		}

		// Save the product to the database
		product := models.Product{
			Name:        name,
			Price:       price,
			Stock:       stock,
			Description: description,
			Image:       handler.Filename,
			IDCategorie: uint(categoryID),
			Down:        "NO",
		}

		err = db.CreateProduct(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to a success page or list of products
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	// If the method is not POST, render the form
	data := map[string]interface{}{
		"Categories": categoriesList,
		"values": map[string]string{
			"name":         "",
			"price":        "",
			"stock":        "",
			"id_categorie": "",
			"description":  "",
			"image":        "",
		},
		"errors": map[string]string{},
	}

	utils.RenderTemplate(w, "templates/back/products/new_product.html", data)
}
