package utils

import (
	"errors"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_web/db"
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"github.com/golang-jwt/jwt/v5"
)

// SecretKey para firmar el token JWT
var SecretKey = []byte("HarwareStore")

// CreateToken genera un token JWT para un usuario dado
func CreateToken(user models.User) (string, error) {
	// Crear el payload del token con los datos del usuario
	payload := models.Claim{
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

// ProcessToken extrae y valida los datos del token JWT
func ProcessToken(tokenString string) (*models.Claim, bool, uint, error) {
	claim := &models.Claim{}

	// Parsear el token con las reclamaciones y la clave secreta
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return claim, false, 0, err
	}

	if token.Valid {
		// Verificar si el usuario existe en la base de datos
		found, _, err := db.CheckExistUser(claim.Email)
		if err != nil {
			return claim, false, 0, err
		}
		if found {
			return claim, true, claim.ID, nil
		}
		return claim, false, 0, errors.New("usuario no encontrado")
	}

	return claim, false, 0, errors.New("token inválido")
}
