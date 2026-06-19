package model

import (
	"time"

	"github.com/google/uuid"
)

type CVDocument struct {
	ID            uuid.UUID `json:"id"`
	CVProfileID   uuid.UUID `json:"cv_profile_id"`
	OriginalFile  []byte    `json:"-"`
	FileType      string    `json:"file_type"`
	ExtractedText string    `json:"extracted_text"`
	UploadedAt    time.Time `json:"uploaded_at"`
}
