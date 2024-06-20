package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/golang-jwt/jwt/v5"
)

// SecretKey para firmar el token JWT
var SecretKey = []byte("HarwareStore")

// Claim define los datos que se almacenarán en el token JWT
type Claim struct {
	Email    string `json:"email"`
	RolID    uint   `json:"rolId"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	ID       uint   `json:"id"`
	jwt.RegisteredClaims
}

// CreateToken genera un token JWT para un usuario dado
func CreateToken(user models.User) (string, error) {
	// Crear el payload del token con los datos del usuario
	payload := Claim{
		Email:    user.Email,
		RolID:    user.IDProfile,
		Name:     user.Name,
		LastName: user.LastName,
		ID:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// Crear un nuevo token con el método de firma HS256 y el payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Variables globales para almacenar datos de los endpoints
var (
	Email  string
	IDUser uint
	RolID  uint
)

// ProcessToken extrae y valida los datos del token JWT
func ProcessToken(tokenString string) (*Claim, bool, uint, error) {
	claim := &Claim{}

	// Dividir el token para obtener solo el valor sin el prefijo "Bearer"
	splitToken := strings.Split(tokenString, "Bearer")
	if len(splitToken) != 2 {
		return claim, false, 0, errors.New("formato de token inválido")
	}

	tokenString = strings.TrimSpace(splitToken[1])

	// Parsear el token con las reclamaciones y la clave secreta
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return claim, false, 0, err
	}

	if token.Valid {
		// Verificar si el usuario existe en la base de datos
		found, err := db.CheckExistUser(claim.Email)
		if err != nil {
			return claim, false, 0, err
		}
		if found {
			Email = claim.Email
			IDUser = claim.ID
			RolID = claim.RolID
			return claim, true, IDUser, nil
		}
		return claim, false, 0, errors.New("usuario no encontrado")
	}

	return claim, false, 0, errors.New("token inválido")
}
