package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type CVProfile struct {
	ID         uuid.UUID       `json:"id"`
	UserID     uuid.UUID       `json:"user_id"`
	FullName   string          `json:"full_name"`
	Headline   string          `json:"headline"`
	Summary    string          `json:"summary"`
	Experience json.RawMessage `json:"experience"`
	Education  json.RawMessage `json:"education"`
	Skills     json.RawMessage `json:"skills"`
	Languages  json.RawMessage `json:"languages"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
