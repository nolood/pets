package jwtclaims

import (
	jwtgo "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwtgo.StandardClaims
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}
