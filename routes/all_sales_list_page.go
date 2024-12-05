package routes

import (
	"net/http"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

type SaleWithUser struct {
	ID         int32     `json:"id"`
	IDUser     int32     `json:"id_user"`
	TotalPrice float64   `json:"total_price"`
	Date       time.Time `json:"date"`
	UserName   string    `json:"user_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
}

func AllSalesListPage(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)

	results, err := db.AllSales()
	if err != nil {
		http.Error(w, "Error obteniendo las ventas", http.StatusInternalServerError)
		return
	}

	sales, err := ConvertToSalesWithUser(results)
	if err != nil {
		http.Error(w, "Error convirtiendo las ventas", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Titulo":    "Mis Pedidos",
		"IDProfile": userData.RolID,
		"Sales":     sales,
	}

	utils.RenderTemplate(w, "templates/back/sales/list_sales_admin.html", data)
}

func ConvertToSalesWithUser(results []map[string]interface{}) ([]SaleWithUser, error) {
	var sales []SaleWithUser

	for _, result := range results {
		sale := SaleWithUser{
			ID:         result["id"].(int32),
			IDUser:     result["id_user"].(int32),
			TotalPrice: result["total_price"].(float64),
			Date:       result["date"].(time.Time),
			UserName:   result["user_name"].(string),
			LastName:   result["last_name"].(string),
			Email:      result["email"].(string),
		}
		sales = append(sales, sale)
	}
	return sales, nil
}
