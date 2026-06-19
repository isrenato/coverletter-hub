package model

import (
	"time"

	"github.com/google/uuid"
)

type CoverLetter struct {
	ID            uuid.UUID  `json:"id"`
	VacancyID     uuid.UUID  `json:"vacancy_id"`
	CVProfileID   uuid.UUID  `json:"cv_profile_id"`
	GeneratedText string     `json:"generated_text"`
	EditedText    string     `json:"edited_text"`
	Status        string     `json:"status"`
	GeneratedAt   time.Time  `json:"generated_at"`
	ApprovedAt    *time.Time `json:"approved_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

const (
	CoverLetterStatusDraft    = "draft"
	CoverLetterStatusApproved = "approved"
	CoverLetterStatusRejected = "rejected"
)
