package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"cyberpets/pets/internal/config"
	"cyberpets/pets/internal/domain/models"
	"cyberpets/pets/internal/domain/telegram"
	"cyberpets/pets/internal/lib/jwt"
	"cyberpets/pets/internal/repositories/user"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Service interface {
	Validate(ctx context.Context, data telegram.WebAppData) (string, error)
}

type userData struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	LastName     string `json:"last_name"`
	FirstName    string `json:"first_name"`
	LanguageCode string `json:"language_code"`
}

type authService struct {
	repo user.Repository
	log  *zap.Logger
	cfg  *config.Config
}

func New(log *zap.Logger, repo user.Repository, cfg *config.Config) Service {
	return &authService{repo: repo, log: log, cfg: cfg}
}

func (s *authService) Validate(ctx context.Context, data telegram.WebAppData) (string, error) {
	const op = "service.auth.validate"

	dataCheckString := fmt.Sprintf("auth_date=%d\nquery_id=%s\nuser=%s", data.AuthDate, data.QueryID, data.User)

	secretKey := hmacSHA256([]byte(data.Token), []byte("WebAppData"))

	signature := hmacSHA256([]byte(dataCheckString), secretKey)

	if hex.EncodeToString(signature) != data.Hash || isDataOutdated(data.AuthDate) {
		return "", fmt.Errorf("%s", op)
	}

	var userDataStruct userData
	err := json.Unmarshal([]byte(data.User), &userDataStruct)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var userModel models.User

	// TODO: to think about - refactoring

	userModel.TgID = userDataStruct.ID
	userModel.Username = userDataStruct.Username
	userModel.LastName = userDataStruct.LastName
	userModel.FirstName = userDataStruct.FirstName
	userModel.LanguageCode = userDataStruct.LanguageCode

	newUser, err := s.repo.CreateOrUpdate(ctx, userModel)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	token, err := jwt.GenerateToken(newUser, s.cfg.Secret)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
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
