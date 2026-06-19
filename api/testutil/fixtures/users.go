package fixtures

import (
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
)

var (
	UserJohnID = uuid.MustParse("a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	UserJaneID = uuid.MustParse("b2c3d4e5-f6a7-8901-bcde-f12345678901")

	UserJohn = model.User{
		ID:           UserJohnID,
		LinkedInID:   "linkedin-john-123",
		Email:        "john@example.com",
		Name:         "John Doe",
		AccessToken:  "access-token-john",
		RefreshToken: "refresh-token-john",
		CreatedAt:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	UserJane = model.User{
		ID:           UserJaneID,
		LinkedInID:   "linkedin-jane-456",
		Email:        "jane@example.com",
		Name:         "Jane Smith",
		AccessToken:  "access-token-jane",
		RefreshToken: "refresh-token-jane",
		CreatedAt:    time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
	}
)
