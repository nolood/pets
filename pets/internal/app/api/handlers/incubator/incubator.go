package incubator

import (
	"cyberpets/pets/internal/app/api/services/incubator"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

var (
	ErrClear   = "err-incubator-clear"
	ErrOpenEgg = "err-incubator-open-eeg"
	ErrSetEgg  = "err-incubator-set-egg"
	ErrGet     = "err-incubator-get"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	SetEgg(w http.ResponseWriter, r *http.Request)
	Clear(w http.ResponseWriter, r *http.Request)
	OpenEgg(w http.ResponseWriter, r *http.Request)
}

type incubatorHandler struct {
	s   incubator.Service
	log *zap.Logger
}

func New(log *zap.Logger, s incubator.Service) Handler {
	return &incubatorHandler{s: s, log: log}
}

// Clear - убрать яйцо из инкубатора

func (h *incubatorHandler) Clear(w http.ResponseWriter, r *http.Request) {

	incubatorInst, err := h.s.Clear(r.Context())
	if err != nil {
		http.Error(w, ErrClear, http.StatusInternalServerError)
		h.log.Error("Cant clear incubator", zap.Error(err))
		return
	}

	response, err := json.Marshal(incubatorInst)
	if err != nil {
		http.Error(w, ErrClear, http.StatusInternalServerError)
		h.log.Error("Cant marshal incubator", zap.Error(err))
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, ErrClear, http.StatusInternalServerError)
		h.log.Error("Cant write incubator", zap.Error(err))
	}

}

// OpenEgg - открыть яйцо в инкубаторе

func (h *incubatorHandler) OpenEgg(w http.ResponseWriter, r *http.Request) {

	eggIdStr := chi.URLParam(r, "eggId")

	eggId, err := strconv.ParseUint(eggIdStr, 10, 64)
	if err != nil {
		http.Error(w, ErrOpenEgg, http.StatusBadRequest)
		h.log.Error("Cant parse egg id", zap.Error(err))
		return
	}

	pet, err := h.s.OpenEgg(r.Context(), eggId)
	if err != nil {
		http.Error(w, ErrOpenEgg, http.StatusInternalServerError)
		h.log.Error("Cant open egg", zap.Error(err))
		return
	}

	response, err := json.Marshal(pet)
	if err != nil {
		http.Error(w, ErrOpenEgg, http.StatusInternalServerError)
		h.log.Error("Cant marshal incubator", zap.Error(err))
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, ErrOpenEgg, http.StatusInternalServerError)
		h.log.Error("Cant write incubator", zap.Error(err))
	}
}

// SetEgg - установить яйцо в инкубаторе

func (h *incubatorHandler) SetEgg(w http.ResponseWriter, r *http.Request) {

	eggIdStr := chi.URLParam(r, "eggId")

	eggId, err := strconv.ParseUint(eggIdStr, 10, 64)

	incubatorInst, err := h.s.SetEgg(r.Context(), eggId)
	if err != nil {
		http.Error(w, ErrSetEgg, http.StatusInternalServerError)
		h.log.Error("Cant set incubator", zap.Error(err))
		return
	}

	response, err := json.Marshal(incubatorInst)
	if err != nil {
		http.Error(w, ErrSetEgg, http.StatusInternalServerError)
		h.log.Error("Cant marshal incubator", zap.Error(err))
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, ErrSetEgg, http.StatusInternalServerError)
		h.log.Error("Cant write incubator", zap.Error(err))
	}
}

// Get - получить инкубатор

func (h *incubatorHandler) Get(w http.ResponseWriter, r *http.Request) {
	incubatorInst, err := h.s.Get(r.Context())
	if err != nil {
		http.Error(w, ErrGet, http.StatusInternalServerError)
		h.log.Error("Cant get incubator", zap.Error(err))
		return
	}

	response, err := json.Marshal(incubatorInst)
	if err != nil {
		http.Error(w, ErrGet, http.StatusInternalServerError)
		h.log.Error("Cant marshal incubator", zap.Error(err))
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, ErrGet, http.StatusInternalServerError)
		h.log.Error("Cant write incubator", zap.Error(err))
	}
}
