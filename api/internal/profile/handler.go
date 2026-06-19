package profile

import (
	"encoding/json"
	"io"
	"net/http"

	"bitbucket.org/irenato/coverletter-hub/api/internal/auth"
	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(auth.UserIDFromContext(r.Context()))
	if err != nil {
		http.Error(w, "invalid user", http.StatusUnauthorized)
		return
	}

	p, err := h.service.Get(r.Context(), userID)
	if err != nil {
		http.Error(w, "profile not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(auth.UserIDFromContext(r.Context()))
	if err != nil {
		http.Error(w, "invalid user", http.StatusUnauthorized)
		return
	}

	var input model.CVProfile
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.service.Update(r.Context(), userID, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(auth.UserIDFromContext(r.Context()))
	if err != nil {
		http.Error(w, "invalid user", http.StatusUnauthorized)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "file too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("cv")
	if err != nil {
		http.Error(w, "missing cv file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read file", http.StatusInternalServerError)
		return
	}

	fileType := "pdf"
	if ct := header.Header.Get("Content-Type"); ct == "application/vnd.openxmlformats-officedocument.wordprocessingml.document" {
		fileType = "docx"
	}

	result, err := h.service.Upload(r.Context(), userID, content, fileType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
