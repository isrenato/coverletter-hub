package fixtures

import (
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
)

var (
	VacancyBackendID  = uuid.MustParse("d4e5f6a7-b8c9-0123-defa-234567890123")
	VacancyFrontendID = uuid.MustParse("e5f6a7b8-c9d0-1234-efab-345678901234")

	VacancyBackend = model.Vacancy{
		ID:          VacancyBackendID,
		UserID:      UserJohnID,
		Title:       "Backend Engineer",
		Company:     "StartupCo",
		Description: "We are looking for a backend engineer proficient in Go and PostgreSQL.",
		Location:    "Amsterdam, NL",
		LinkedInURL: "https://linkedin.com/jobs/12345",
		Source:      "manual",
		CreatedAt:   time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC),
	}

	VacancyFrontend = model.Vacancy{
		ID:          VacancyFrontendID,
		UserID:      UserJohnID,
		Title:       "Frontend Developer",
		Company:     "DesignLab",
		Description: "React/Vue developer needed for our design tools platform.",
		Location:    "Remote",
		Source:      "manual",
		CreatedAt:   time.Date(2026, 6, 10, 0, 0, 0, 0, time.UTC),
	}
)
