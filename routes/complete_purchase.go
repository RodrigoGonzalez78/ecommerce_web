package routes

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func CompletePurchase(w http.ResponseWriter, r *http.Request) {

	userData, _ := r.Context().Value("userData").(*models.Claim)

	user, err := db.GetUser(userData.ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}
	if user.IDAddress == nil {
		http.Redirect(w, r, "/user-profile", http.StatusSeeOther)
		return
	}

	// Leer la cookie del carrito
	cartCookie, err := r.Cookie("cart")

	if err != nil {
		http.Error(w, "Carrito no encontrado", http.StatusNotFound)
		return
	}

	// Decodificar el carrito
	cartBytes, err := base64.URLEncoding.DecodeString(cartCookie.Value)
	if err != nil {
		http.Error(w, "Error al decodificar el carrito", http.StatusInternalServerError)
		return
	}

	var cart []models.CartItem
	err = json.Unmarshal(cartBytes, &cart)
	if err != nil {
		http.Error(w, "Error al deserializar el carrito", http.StatusInternalServerError)
		return
	}

	// Calcular el precio total de la venta
	var totalPrice float64
	for _, item := range cart {
		product, err := db.GetProduct(uint(item.ID))
		if err != nil {
			http.Error(w, "Producto no encontrado", http.StatusInternalServerError)
			return
		}
		totalPrice += product.Price * float64(item.Qty)
	}

	// Crear la venta
	sale := models.Sale{
		IDUser:     userData.ID,
		TotalPrice: totalPrice,
		Date:       time.Now(),
	}

	saleCreated, err := db.CreateSale(sale)
	if err != nil {
		http.Error(w, "Error al crear la venta", http.StatusInternalServerError)
		return
	}

	// Crear los detalles de la venta y actualizar el stock
	for _, item := range cart {
		product, err := db.GetProduct(uint(item.ID))

		if err != nil {
			http.Error(w, "Producto no encontrado", http.StatusInternalServerError)
			return
		}

		saleDetails := models.SaleDetails{
			IDSale:    saleCreated.ID,
			Count:     item.Qty,
			Price:     product.Price,
			IDProduct: product.ID,
		}

		err = db.CreateSaleDetails(saleDetails)
		if err != nil {
			http.Error(w, "Error al crear los detalles de la venta", http.StatusInternalServerError)
			return
		}

		// Actualizar el stock del producto
		newStock := product.Stock - item.Qty
		err = db.UpdateProduct(product.ID, map[string]interface{}{"stock": newStock})
		if err != nil {
			http.Error(w, "Error al actualizar el stock del producto", http.StatusInternalServerError)
			return
		}
	}

	// Vaciar el carrito
	expiration := time.Now().Add(-24 * time.Hour)
	cartCookie = &http.Cookie{
		Name:     "cart",
		Value:    "",
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, cartCookie)

	http.Redirect(w, r, "/my-orders", http.StatusSeeOther)
}
