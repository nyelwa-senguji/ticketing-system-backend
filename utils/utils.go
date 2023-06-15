package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func LoadEnviromentalVariables(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(".env file is missing")
	}
	return os.Getenv(key)
}

func GenerateJWTToken(id string) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    id,
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)}, // 1 Day
	})

	token, err := claims.SignedString([]byte(LoadEnviromentalVariables("SECRET_KEY")))
	if err != nil {
		return "Could not login"
	}
	return token
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
