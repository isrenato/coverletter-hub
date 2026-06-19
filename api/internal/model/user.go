package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	LinkedInID   string    `json:"linkedin_id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	AccessToken  string    `json:"-"`
	RefreshToken string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
