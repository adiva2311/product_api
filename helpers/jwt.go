package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
)

type jwtCustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(id int, username string, role string) (string, error) {
	// Set custom claims
	customClaims := &jwtCustomClaims{
		ID:       uint(id),
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Error("Unable to generate token")
		return "", err
	}

	return signedToken, nil
}

func GetSecretKey() []byte {
	return secretKey
}
