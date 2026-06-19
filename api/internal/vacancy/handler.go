package vacancy

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/auth"
	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(auth.UserIDFromContext(r.Context()))
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var input struct {
		Title       string `json:"title"`
		Company     string `json:"company"`
		Description string `json:"description"`
		Location    string `json:"location"`
		LinkedInURL string `json:"linkedin_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	v := &model.Vacancy{
		ID:          uuid.New(),
		UserID:      userID,
		Title:       input.Title,
		Company:     input.Company,
		Description: input.Description,
		Location:    input.Location,
		LinkedInURL: input.LinkedInURL,
		Source:      "manual",
		CreatedAt:   time.Now(),
	}

	if err := h.repo.Create(r.Context(), v); err != nil {
		http.Error(w, "failed to create vacancy", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid vacancy id", http.StatusBadRequest)
		return
	}

	v, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "vacancy not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) HandleList(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(auth.UserIDFromContext(r.Context()))
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if limit <= 0 {
		limit = 20
	}

	items, total, err := h.repo.List(r.Context(), userID, ListOptions{Limit: limit, Offset: offset})
	if err != nil {
		http.Error(w, "failed to list vacancies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": items,
		"total": total,
	})
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid vacancy id", http.StatusBadRequest)
		return
	}

	existing, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "vacancy not found", http.StatusNotFound)
		return
	}

	var input struct {
		Title       string `json:"title"`
		Company     string `json:"company"`
		Description string `json:"description"`
		Location    string `json:"location"`
		LinkedInURL string `json:"linkedin_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	existing.Title = input.Title
	existing.Company = input.Company
	existing.Description = input.Description
	existing.Location = input.Location
	existing.LinkedInURL = input.LinkedInURL

	if err := h.repo.Update(r.Context(), existing); err != nil {
		http.Error(w, "failed to update vacancy", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existing)
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid vacancy id", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(r.Context(), id); err != nil {
		http.Error(w, "vacancy not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
