package fixtures

import (
	"encoding/json"
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
)

var (
	CVProfileJohnID = uuid.MustParse("c3d4e5f6-a7b8-9012-cdef-123456789012")

	CVProfileJohn = model.CVProfile{
		ID:         CVProfileJohnID,
		UserID:     UserJohnID,
		FullName:   "John Doe",
		Headline:   "Senior Software Engineer",
		Summary:    "10 years of experience building web applications.",
		Experience: json.RawMessage(`[{"title":"Senior Engineer","company":"TechCorp","start":"2020-01","end":"present"}]`),
		Education:  json.RawMessage(`[{"degree":"BSc Computer Science","school":"MIT","year":"2015"}]`),
		Skills:     json.RawMessage(`["Go","TypeScript","PostgreSQL","Docker"]`),
		Languages:  json.RawMessage(`["English","Dutch"]`),
		CreatedAt:  time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:  time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
	}
)
