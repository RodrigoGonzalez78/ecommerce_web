package models

type CartItem struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
