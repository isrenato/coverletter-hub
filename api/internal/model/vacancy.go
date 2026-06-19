package model

import (
	"time"

	"github.com/google/uuid"
)

type Vacancy struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	LinkedInURL string    `json:"linkedin_url,omitempty"`
	Source      string    `json:"source"`
	CreatedAt   time.Time `json:"created_at"`
}
