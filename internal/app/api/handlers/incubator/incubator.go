package incubator

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"pets/internal/app/api/services/incubator"
	"strconv"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	SetEgg(w http.ResponseWriter, r *http.Request)
	RemoveEgg(w http.ResponseWriter, r *http.Request)
	OpenEgg(w http.ResponseWriter, r *http.Request)
}

type incubatorHandler struct {
	s   incubator.Service
	log *zap.Logger
}

func New(log *zap.Logger, s incubator.Service) Handler {
	return &incubatorHandler{s: s, log: log}
}

func (h *incubatorHandler) RemoveEgg(w http.ResponseWriter, r *http.Request) {
}

func (h *incubatorHandler) OpenEgg(w http.ResponseWriter, r *http.Request) {

	eggIdStr := chi.URLParam(r, "eggId")

	eggId, err := strconv.ParseUint(eggIdStr, 10, 64)
	if err != nil {
		http.Error(w, "err-egg-id", http.StatusBadRequest)
		h.log.Error("Cant parse egg id", zap.Error(err))
		return
	}

	pet, err := h.s.OpenEgg(r.Context(), eggId)
	if err != nil {
		http.Error(w, "err-egg-open", http.StatusInternalServerError)
		h.log.Error("Cant open egg", zap.Error(err))
		return
	}

	response, err := json.Marshal(pet)
	if err != nil {
		http.Error(w, "err-egg-marshal", http.StatusInternalServerError)
		h.log.Error("Cant marshal incubator", zap.Error(err))
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "err-egg-write", http.StatusInternalServerError)
		h.log.Error("Cant write incubator", zap.Error(err))
	}
}

func (h *incubatorHandler) SetEgg(w http.ResponseWriter, r *http.Request) {
}

func (h *incubatorHandler) Get(w http.ResponseWriter, r *http.Request) {
	incubatorInst, err := h.s.Get(r.Context())
	if err != nil {
		http.Error(w, "err-incubator-get", http.StatusInternalServerError)
		h.log.Error("Cant get incubator", zap.Error(err))
		return
	}

	response, err := json.Marshal(incubatorInst)
	if err != nil {
		http.Error(w, "err-incubator-marshal", http.StatusInternalServerError)
		h.log.Error("Cant marshal incubator", zap.Error(err))
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "err-incubator-write", http.StatusInternalServerError)
		h.log.Error("Cant write incubator", zap.Error(err))
	}
}
