package auth

import (
	"context"
	jwtclaims "cyberpets/jwt-claims"
	"net/http"
	"strings"

	jwtgo "github.com/dgrijalva/jwt-go"
)

func Middleware(secret string, env string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if env == "local" {
				ctx := context.WithValue(r.Context(), "user_id", 1)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			authorizationHeader := r.Header.Get("Authorization")

			if authorizationHeader == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authorizationHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			var myClaims jwtclaims.Claims
			token, err := jwtgo.ParseWithClaims(tokenString, &myClaims, func(token *jwtgo.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			userID := myClaims.Id

			ctx := context.WithValue(r.Context(), "user_id", userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
