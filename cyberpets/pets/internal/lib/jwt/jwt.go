package jwt

import (
	jwtclaims "cyberpets/jwt-claims"
	"cyberpets/pets/internal/domain/models"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

func GenerateToken(user models.User, secret string) (string, error) {
	claims := jwtclaims.Claims{
		Username: user.Username,
		ID:       user.ID,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
