package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

func MyOrders(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)
	sales, _ := db.UserSales(userData.ID)

	data := map[string]interface{}{
		"Titulo":    "Mis Pedidos",
		"IDProfile": userData.RolID,
		"Sales":     sales,
	}

	utils.RenderTemplate(w, "templates/back/sales/list_sales.html", data)
}
