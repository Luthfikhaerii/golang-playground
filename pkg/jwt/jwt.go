package jwt

import (
	"golang-playground/internal/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload *model.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    payload.ID,
		"name":  payload.Name,
		"email": payload.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
