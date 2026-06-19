package coverletter_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/coverletter"
	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockCLRepo struct {
	letters map[uuid.UUID]*model.CoverLetter
}

func newMockCLRepo() *mockCLRepo {
	return &mockCLRepo{letters: make(map[uuid.UUID]*model.CoverLetter)}
}

func (m *mockCLRepo) Create(_ context.Context, cl *model.CoverLetter) error {
	m.letters[cl.ID] = cl
	return nil
}

func (m *mockCLRepo) GetByID(_ context.Context, id uuid.UUID) (*model.CoverLetter, error) {
	if cl, ok := m.letters[id]; ok {
		return cl, nil
	}
	return nil, coverletter.ErrNotFound
}

func (m *mockCLRepo) List(_ context.Context, _ uuid.UUID, _ coverletter.ListOptions) ([]model.CoverLetter, int, error) {
	var items []model.CoverLetter
	for _, cl := range m.letters {
		items = append(items, *cl)
	}
	return items, len(items), nil
}

func (m *mockCLRepo) Update(_ context.Context, cl *model.CoverLetter) error {
	m.letters[cl.ID] = cl
	return nil
}

func (m *mockCLRepo) GetByVacancyID(_ context.Context, vacancyID uuid.UUID) ([]model.CoverLetter, error) {
	var items []model.CoverLetter
	for _, cl := range m.letters {
		if cl.VacancyID == vacancyID {
			items = append(items, *cl)
		}
	}
	return items, nil
}

type mockVacancyRepo struct {
	vacancies map[uuid.UUID]*model.Vacancy
}

func newMockVacancyRepo() *mockVacancyRepo {
	r := &mockVacancyRepo{vacancies: make(map[uuid.UUID]*model.Vacancy)}
	v := fixtures.VacancyBackend
	r.vacancies[v.ID] = &v
	return r
}

func (m *mockVacancyRepo) GetByID(_ context.Context, id uuid.UUID) (*model.Vacancy, error) {
	if v, ok := m.vacancies[id]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("not found")
}

type mockProfileRepo2 struct{}

func (m *mockProfileRepo2) GetByUserID(_ context.Context, _ uuid.UUID) (*model.CVProfile, error) {
	p := fixtures.CVProfileJohn
	return &p, nil
}

type mockLLM2 struct{}

func (m *mockLLM2) Generate(_ context.Context, _, _ string) (string, error) {
	return "Dear Hiring Manager,\n\nI am excited to apply...", nil
}

func TestServiceGenerate_CreatesDraft(t *testing.T) {
	clRepo := newMockCLRepo()
	svc := coverletter.NewService(clRepo, newMockVacancyRepo(), &mockProfileRepo2{}, &mockLLM2{})

	result, err := svc.Generate(context.Background(), fixtures.UserJohnID, fixtures.VacancyBackendID)
	require.NoError(t, err)
	assert.Equal(t, model.CoverLetterStatusDraft, result.Status)
	assert.Contains(t, result.GeneratedText, "Dear Hiring Manager")
	assert.False(t, result.HasWarning)
}

func TestServiceGenerate_WarnsIfRecentApproved(t *testing.T) {
	clRepo := newMockCLRepo()
	recent := fixtures.CoverLetterDraft
	recent.ID = uuid.New()
	recent.Status = model.CoverLetterStatusApproved
	now := time.Now().Add(-3 * 30 * 24 * time.Hour)
	recent.ApprovedAt = &now
	_ = clRepo.Create(context.Background(), &recent)

	svc := coverletter.NewService(clRepo, newMockVacancyRepo(), &mockProfileRepo2{}, &mockLLM2{})

	result, err := svc.Generate(context.Background(), fixtures.UserJohnID, fixtures.VacancyBackendID)
	require.NoError(t, err)
	assert.True(t, result.HasWarning)
}

func TestServiceUpdateStatus_Approve(t *testing.T) {
	clRepo := newMockCLRepo()
	draft := fixtures.CoverLetterDraft
	_ = clRepo.Create(context.Background(), &draft)

	svc := coverletter.NewService(clRepo, newMockVacancyRepo(), &mockProfileRepo2{}, &mockLLM2{})

	result, err := svc.UpdateStatus(context.Background(), fixtures.CoverLetterDraftID, "approved")
	require.NoError(t, err)
	assert.Equal(t, model.CoverLetterStatusApproved, result.Status)
	assert.NotNil(t, result.ApprovedAt)
}
