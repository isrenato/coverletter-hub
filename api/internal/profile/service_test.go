package profile_test

import (
	"context"
	"encoding/json"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"bitbucket.org/irenato/coverletter-hub/api/internal/profile"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockProfileRepo struct {
	profiles map[uuid.UUID]*model.CVProfile
}

func newMockProfileRepo() *mockProfileRepo {
	return &mockProfileRepo{profiles: make(map[uuid.UUID]*model.CVProfile)}
}

func (m *mockProfileRepo) Create(_ context.Context, p *model.CVProfile) error {
	m.profiles[p.UserID] = p
	return nil
}

func (m *mockProfileRepo) GetByUserID(_ context.Context, userID uuid.UUID) (*model.CVProfile, error) {
	if p, ok := m.profiles[userID]; ok {
		return p, nil
	}
	return nil, profile.ErrNotFound
}

func (m *mockProfileRepo) Update(_ context.Context, p *model.CVProfile) error {
	m.profiles[p.UserID] = p
	return nil
}

type mockDocRepo struct {
	docs []model.CVDocument
}

func (m *mockDocRepo) Create(_ context.Context, doc *model.CVDocument) error {
	m.docs = append(m.docs, *doc)
	return nil
}

func (m *mockDocRepo) GetByProfileID(_ context.Context, _ uuid.UUID) ([]model.CVDocument, error) {
	return m.docs, nil
}

type mockParser struct {
	result *model.CVProfile
	err    error
}

func (m *mockParser) ParseCV(_ context.Context, _ []byte, _ string) (*model.CVProfile, error) {
	return m.result, m.err
}

func TestServiceGet_ExistingProfile(t *testing.T) {
	repo := newMockProfileRepo()
	_ = repo.Create(context.Background(), &fixtures.CVProfileJohn)

	svc := profile.NewService(repo, &mockDocRepo{}, &mockParser{})

	got, err := svc.Get(context.Background(), fixtures.UserJohnID)
	require.NoError(t, err)
	assert.Equal(t, fixtures.CVProfileJohn.FullName, got.FullName)
}

func TestServiceUpload_CreatesProfileAndDocument(t *testing.T) {
	repo := newMockProfileRepo()
	docRepo := &mockDocRepo{}
	parsed := &model.CVProfile{
		FullName:   "Parsed Name",
		Headline:   "Parsed Headline",
		Summary:    "Parsed summary",
		Experience: json.RawMessage(`[]`),
		Education:  json.RawMessage(`[]`),
		Skills:     json.RawMessage(`[]`),
		Languages:  json.RawMessage(`[]`),
	}
	parser := &mockParser{result: parsed}

	svc := profile.NewService(repo, docRepo, parser)

	result, err := svc.Upload(context.Background(), fixtures.UserJohnID, []byte("pdf data"), "pdf")
	require.NoError(t, err)
	assert.Equal(t, "Parsed Name", result.FullName)
	require.Len(t, docRepo.docs, 1)
	assert.Equal(t, "pdf", docRepo.docs[0].FileType)
}
