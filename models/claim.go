package models

import "github.com/golang-jwt/jwt/v5"

// Claim define los datos que se almacenarán en el token JWT
type Claim struct {
	Email    string `json:"email"`
	RolID    uint   `json:"rolId"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	ID       uint   `json:"id"`
	jwt.RegisteredClaims
}
