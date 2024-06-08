package jwt

import (
	"cyberpets/pets/internal/domain/models"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	jwtgo.StandardClaims
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func GenerateToken(user models.User, secret string) (string, error) {
	claims := Claims{
		Username: user.Username,
		ID:       user.ID,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
