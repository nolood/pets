package farm

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"pets/internal/app/api/services/farm"
	"strconv"
)

var (
	ErrGet = "err-farm-get"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	SetPet(w http.ResponseWriter, r *http.Request)
	RemovePet(w http.ResponseWriter, r *http.Request)
}

type farmHandler struct {
	service farm.Service
	log     *zap.Logger
}

func New(log *zap.Logger, service farm.Service) Handler {
	return &farmHandler{service: service, log: log}
}

func (h *farmHandler) Get(w http.ResponseWriter, r *http.Request) {
	farmInst, err := h.service.Get(r.Context())
	if err != nil {
		http.Error(w, ErrGet, http.StatusInternalServerError)
		h.log.Error("Cant get farm", zap.Error(err))
		return
	}

	response, err := json.Marshal(farmInst)
	if err != nil {
		http.Error(w, ErrGet, http.StatusInternalServerError)
		h.log.Error("Cant get farm", zap.Error(err))
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, ErrGet, http.StatusInternalServerError)
		h.log.Error("Cant get farm", zap.Error(err))
	}
}

func (h *farmHandler) SetPet(w http.ResponseWriter, r *http.Request) {
	strSlotId := chi.URLParam(r, "slotId")
	strPetId := chi.URLParam(r, "petId")

	slotId, err := strconv.ParseUint(strSlotId, 10, 64)
	if err != nil {
		http.Error(w, ErrGet, http.StatusBadRequest)
		h.log.Error("Cant parse slot id", zap.Error(err))
		return
	}

	petId, err := strconv.ParseUint(strPetId, 10, 64)
	if err != nil {
		http.Error(w, ErrGet, http.StatusBadRequest)
		h.log.Error("Cant parse pet id", zap.Error(err))
		return
	}

	farmInst, err := h.service.SetPet(r.Context(), petId, slotId)
	if err != nil {
		http.Error(w, ErrGet, http.StatusInternalServerError)
		h.log.Error("Cant set pet", zap.Error(err))
		return
	}

	response, err := json.Marshal(farmInst)
	if err != nil {
		http.Error(w, ErrGet, http.StatusInternalServerError)
		h.log.Error("Cant marshal farm", zap.Error(err))
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, ErrGet, http.StatusInternalServerError)
		h.log.Error("Cant write farm", zap.Error(err))
	}

}

func (h *farmHandler) RemovePet(w http.ResponseWriter, r *http.Request) {

}
