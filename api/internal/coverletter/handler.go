package coverletter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bitbucket.org/irenato/coverletter-hub/api/internal/auth"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleGenerate(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(auth.UserIDFromContext(r.Context()))
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	vacancyID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid vacancy id", http.StatusBadRequest)
		return
	}

	result, err := h.service.Generate(r.Context(), userID, vacancyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid cover letter id", http.StatusBadRequest)
		return
	}

	cl, err := h.service.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "cover letter not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cl)
}

func (h *Handler) HandleList(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(auth.UserIDFromContext(r.Context()))
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	status := r.URL.Query().Get("status")
	if limit <= 0 {
		limit = 20
	}

	items, total, err := h.service.repo.List(r.Context(), userID, ListOptions{Limit: limit, Offset: offset, Status: status})
	if err != nil {
		http.Error(w, "failed to list cover letters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": items,
		"total": total,
	})
}

func (h *Handler) HandleUpdateText(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid cover letter id", http.StatusBadRequest)
		return
	}

	var input struct {
		EditedText string `json:"edited_text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.service.UpdateText(r.Context(), id, input.EditedText)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) HandleUpdateStatus(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid cover letter id", http.StatusBadRequest)
		return
	}

	var input struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.service.UpdateStatus(r.Context(), id, input.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
