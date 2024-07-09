package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	jwtclaims "cyberpets/jwt-claims"
	ssov1 "cyberpets/protos/gen/go/sso"
	"encoding/hex"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	Validate(ctx context.Context, req *ssov1.ValidateRequest) bool
	ValidateToken(ctx context.Context, token string) (bool, string)
}

type service struct {
	log     *zap.Logger
	tgToken string
	secret  string
}

func New(log *zap.Logger, token string, secret string) Service {
	return &service{log: log, tgToken: token, secret: secret}
}

func (s *service) Validate(ctx context.Context, req *ssov1.ValidateRequest) bool {
	dataCheckString := fmt.Sprintf("auth_date=%d\nquery_id=%s\nuser=%s", req.AuthDate, req.QueryId, req.User)

	secretKey := hmacSHA256([]byte(s.tgToken), []byte("WebAppData"))

	signature := hmacSHA256([]byte(dataCheckString), secretKey)

	if hex.EncodeToString(signature) != req.Hash || isDataOutdated(req.AuthDate) {
		return false
	}

	return true
}

func (s *service) ValidateToken(ctx context.Context, tokenStr string) (bool, string) {
	var myClaims jwtclaims.Claims
	token, err := jwtgo.ParseWithClaims(tokenStr, &myClaims, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return false, ""
	}

	if !token.Valid {
		return false, ""
	}

	return true, myClaims.Id
}

func hmacSHA256(data []byte, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func isDataOutdated(authDate int64) bool {
	currentTime := time.Now().Unix()
	if currentTime-authDate > 24*3600 {
		return true
	}
	return false
}
