package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, u *model.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByLinkedInID(ctx context.Context, linkedInID string) (*model.User, error)
	Update(ctx context.Context, u *model.User) error
}

type contextKey string

const userIDKey contextKey = "userID"

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func UserIDFromContext(ctx context.Context) string {
	v, _ := ctx.Value(userIDKey).(string)
	return v
}

type AuthHandler struct {
	userRepo    UserRepository
	linkedIn    *LinkedInClient
	jwtSecret   string
	frontendURL string
}

func NewAuthHandler(userRepo UserRepository, linkedIn *LinkedInClient, jwtSecret, frontendURL string) *AuthHandler {
	return &AuthHandler{
		userRepo:    userRepo,
		linkedIn:    linkedIn,
		jwtSecret:   jwtSecret,
		frontendURL: frontendURL,
	}
}

func (h *AuthHandler) HandleLinkedInRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, h.linkedIn.AuthURL(), http.StatusTemporaryRedirect)
}

func (h *AuthHandler) HandleLinkedInCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing code parameter", http.StatusBadRequest)
		return
	}

	tokenResp, err := h.linkedIn.ExchangeCode(r.Context(), code)
	if err != nil {
		log.Printf("linkedin code exchange error: %v", err)
		http.Error(w, "authentication failed", http.StatusInternalServerError)
		return
	}

	profile, err := h.linkedIn.GetProfile(r.Context(), tokenResp.AccessToken)
	if err != nil {
		log.Printf("linkedin profile fetch error: %v", err)
		http.Error(w, "failed to fetch profile", http.StatusInternalServerError)
		return
	}

	user, err := h.userRepo.GetByLinkedInID(r.Context(), profile.Sub)
	if err != nil {
		now := time.Now()
		user = &model.User{
			ID:           uuid.New(),
			LinkedInID:   profile.Sub,
			Email:        profile.Email,
			Name:         profile.Name,
			AccessToken:  tokenResp.AccessToken,
			RefreshToken: tokenResp.RefreshToken,
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		if err := h.userRepo.Create(r.Context(), user); err != nil {
			log.Printf("user create error: %v", err)
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}
	} else {
		user.AccessToken = tokenResp.AccessToken
		user.RefreshToken = tokenResp.RefreshToken
		user.Name = profile.Name
		user.Email = profile.Email
		if err := h.userRepo.Update(r.Context(), user); err != nil {
			log.Printf("user update error: %v", err)
			http.Error(w, "failed to update user", http.StatusInternalServerError)
			return
		}
	}

	jwt, err := GenerateToken(*user, h.jwtSecret)
	if err != nil {
		log.Printf("jwt generation error: %v", err)
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, h.frontendURL+"/#token="+jwt, http.StatusTemporaryRedirect)
}

func (h *AuthHandler) HandleMe(w http.ResponseWriter, r *http.Request) {
	userID := UserIDFromContext(r.Context())
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetByID(r.Context(), uid)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
