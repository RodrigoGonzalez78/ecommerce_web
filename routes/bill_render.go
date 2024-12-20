package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
)

func BillRender(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "ID de la venta no proporcionado", http.StatusBadRequest)
		return
	}

	saleID, err := strconv.Atoi(idParam)
	if err != nil || saleID <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	sale, err := db.FindSale(uint(saleID))
	if err != nil {
		http.Error(w, fmt.Sprintf("Venta con ID %d no encontrada", saleID), http.StatusNotFound)
		return
	}

	user, err := db.GetUser(sale.IDUser)
	if err != nil {
		http.Error(w, "Error al obtener el usuario", http.StatusInternalServerError)
		return
	}

	address, err := db.GetAddress(*user.IDAddress)
	if err != nil {
		http.Error(w, "Error al obtener la dirección", http.StatusInternalServerError)
		return
	}

	saleDetails, err := db.GetSaleDetailsByIdAndProductName(sale.ID)
	if err != nil {
		http.Error(w, "Error al obtener los detalles de la venta", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Sale":        *sale,
		"User":        *user,
		"Address":     *address,
		"SaleDetails": saleDetails,
	}

	tmpl := template.Must(template.ParseFiles("templates/back/sales/bill.html"))

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
	}
}
