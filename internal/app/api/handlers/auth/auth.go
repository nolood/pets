package auth

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"pets/internal/app/api/services/auth"
	"pets/internal/domain/telegram"
)

var (
	ErrValidate = "err-data-validate"
)

type Handler interface {
	Validate(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	service auth.Service
	log     *zap.Logger
	token   string
}

func New(log *zap.Logger, service auth.Service, token string) Handler {
	return &authHandler{
		log:     log,
		service: service,
		token:   token,
	}
}

func (h *authHandler) Validate(w http.ResponseWriter, r *http.Request) {

	var webAppData telegram.WebAppData

	userStr := chi.URLParam(r, "user")
	webAppData.User = userStr
	webAppData.Token = h.token

	err := json.NewDecoder(r.Body).Decode(&webAppData)
	if err != nil {
		h.log.Error("Cant decode", zap.Error(err))
		http.Error(w, ErrValidate, http.StatusBadRequest)
		return
	}

	token, err := h.service.Validate(r.Context(), webAppData)
	if err != nil {
		h.log.Error("Cant validate", zap.Error(err))
		http.Error(w, ErrValidate, http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte(token))
	if err != nil {
		h.log.Error("Cant write", zap.Error(err))
		http.Error(w, ErrValidate, http.StatusInternalServerError)
	}

}
