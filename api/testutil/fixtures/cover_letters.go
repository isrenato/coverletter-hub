package fixtures

import (
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
)

var (
	CoverLetterDraftID    = uuid.MustParse("f6a7b8c9-d0e1-2345-fabc-456789012345")
	CoverLetterApprovedID = uuid.MustParse("a7b8c9d0-e1f2-3456-abcd-567890123456")

	CoverLetterDraft = model.CoverLetter{
		ID:            CoverLetterDraftID,
		VacancyID:     VacancyBackendID,
		CVProfileID:   CVProfileJohnID,
		GeneratedText: "Dear Hiring Manager,\n\nI am writing to express my interest in the Backend Engineer position at StartupCo...",
		EditedText:    "",
		Status:        model.CoverLetterStatusDraft,
		GeneratedAt:   time.Date(2026, 6, 1, 12, 0, 0, 0, time.UTC),
		ApprovedAt:    nil,
		CreatedAt:     time.Date(2026, 6, 1, 12, 0, 0, 0, time.UTC),
		UpdatedAt:     time.Date(2026, 6, 1, 12, 0, 0, 0, time.UTC),
	}

	approvedAt        = time.Date(2026, 1, 15, 10, 0, 0, 0, time.UTC)
	CoverLetterApproved = model.CoverLetter{
		ID:            CoverLetterApprovedID,
		VacancyID:     VacancyFrontendID,
		CVProfileID:   CVProfileJohnID,
		GeneratedText: "Dear Hiring Manager,\n\nI am excited to apply for the Frontend Developer role...",
		EditedText:    "Dear Hiring Manager,\n\nI am thrilled to apply for the Frontend Developer role...",
		Status:        model.CoverLetterStatusApproved,
		GeneratedAt:   time.Date(2026, 1, 15, 9, 0, 0, 0, time.UTC),
		ApprovedAt:    &approvedAt,
		CreatedAt:     time.Date(2026, 1, 15, 9, 0, 0, 0, time.UTC),
		UpdatedAt:     time.Date(2026, 1, 15, 10, 0, 0, 0, time.UTC),
	}
)
