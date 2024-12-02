package routes

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/RodrigoGonzalez78/ecommerce_web/utils"
)

type ProductCartItem struct {
	ID         int
	Name       string
	Image      string
	Amount     int
	Price      float64
	TotalPrice float64
}

func MyCartPage(w http.ResponseWriter, r *http.Request) {
	userData, _ := r.Context().Value("userData").(*models.Claim)

	cartCookie, err := r.Cookie("cart")
	var cartItems []models.CartItem

	if err == nil {

		cartBytes, err := base64.URLEncoding.DecodeString(cartCookie.Value)

		if err != nil {
			http.Error(w, "Error al decodificar el carrito", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(cartBytes, &cartItems)

		if err != nil {
			http.Error(w, "Error al deserializar el carrito", http.StatusInternalServerError)
			return
		}
	}

	products := []ProductCartItem{}

	for _, v := range cartItems {
		prod, _ := db.GetProduct(uint(v.ID))

		products = append(products, ProductCartItem{
			ID:         v.ID,
			Name:       prod.Name,
			Image:      prod.Image,
			Amount:     v.Qty,
			Price:      prod.Price,
			TotalPrice: prod.Price * float64(v.Qty),
		})
	}

	data := map[string]interface{}{
		"Titulo":    "Mi Carrito",
		"IDProfile": userData.RolID,
		"Products":  products,
	}

	utils.RenderTemplate(w, "templates/back/cart/cart_list.html", data)
}
