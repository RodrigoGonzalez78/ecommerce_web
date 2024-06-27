package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
	"github.com/google/uuid"
)

func EditProduct(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := db.GetProduct(uint(id))
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	categoriesList, _ := db.GetCategories()

	if r.Method == http.MethodPost {

		product.Name = r.FormValue("name")
		product.Price, _ = strconv.ParseFloat(r.FormValue("price"), 64)
		product.Stock, _ = strconv.Atoi(r.FormValue("stock"))
		idCategorie, _ := strconv.Atoi(r.FormValue("id_categorie"))
		product.IDCategorie = uint(idCategorie)
		product.Description = r.FormValue("description")

		file, handler, err := r.FormFile("image")

		if err == nil {
			defer file.Close()

			extension := filepath.Ext(handler.Filename)
			randomFileName := uuid.New().String() + extension

			uploadDir := "assets/uploads"
			if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
				fmt.Println("Error creating directory:", err)
				http.Error(w, "Unable to create directory for uploads", http.StatusInternalServerError)
				return
			}

			imagePath := filepath.Join(uploadDir, randomFileName)
			fmt.Println("Saving file to:", imagePath)

			dst, err := os.Create(imagePath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, file)
			if err != nil {
				fmt.Println("Error saving file:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			product.Image = randomFileName
		}

		data := map[string]interface{}{
			"name":         product.Name,
			"price":        product.Price,
			"stock":        product.Stock,
			"id_categorie": product.IDCategorie,
			"description":  product.Description,
			"image":        product.Image,
		}

		err = db.UpdateProduct(product.ID, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/products-page", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Titulo":     "Editar Prouducto",
		"Product":    product,
		"Categories": categoriesList,
		"values": map[string]string{
			"name":         product.Name,
			"price":        fmt.Sprintf("%f", product.Price),
			"stock":        fmt.Sprintf("%d", product.Stock),
			"id_categorie": fmt.Sprintf("%d", product.IDCategorie),
			"description":  product.Description,
		},
		"errors": map[string]string{},
	}

	utils.RenderTemplate(w, "templates/back/products/edit_product.html", data)
}
