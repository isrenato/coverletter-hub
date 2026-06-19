package coverletter_test

import (
	"context"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/coverletter"
	"bitbucket.org/irenato/coverletter-hub/api/testutil"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupCLRepo(t *testing.T) (*coverletter.PostgresRepository, *testutil.TestDB) {
	tdb := testutil.NewTestDB(t)
	testutil.MigrateTestDB(t, tdb.Pool)
	testutil.SeedUser(t, tdb.Pool, fixtures.UserJohn)
	testutil.SeedCVProfile(t, tdb.Pool, fixtures.CVProfileJohn)
	testutil.SeedVacancy(t, tdb.Pool, fixtures.VacancyBackend)
	testutil.SeedVacancy(t, tdb.Pool, fixtures.VacancyFrontend)
	return coverletter.NewPostgresRepository(tdb.Pool), tdb
}

func TestCLCreate_And_GetByID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupCLRepo(t)
	ctx := context.Background()

	err := repo.Create(ctx, &fixtures.CoverLetterDraft)
	require.NoError(t, err)

	got, err := repo.GetByID(ctx, fixtures.CoverLetterDraftID)
	require.NoError(t, err)
	assert.Equal(t, "draft", got.Status)
	assert.Equal(t, fixtures.VacancyBackendID, got.VacancyID)
}

func TestCLGetByVacancyID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupCLRepo(t)
	ctx := context.Background()

	_ = repo.Create(ctx, &fixtures.CoverLetterDraft)

	items, err := repo.GetByVacancyID(ctx, fixtures.VacancyBackendID)
	require.NoError(t, err)
	require.Len(t, items, 1)
}

func TestCLList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupCLRepo(t)
	ctx := context.Background()

	_ = repo.Create(ctx, &fixtures.CoverLetterDraft)
	_ = repo.Create(ctx, &fixtures.CoverLetterApproved)

	items, total, err := repo.List(ctx, fixtures.UserJohnID, coverletter.ListOptions{Limit: 10, Offset: 0})
	require.NoError(t, err)
	assert.Equal(t, 2, total)
	assert.Len(t, items, 2)
}

func TestCLUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupCLRepo(t)
	ctx := context.Background()

	_ = repo.Create(ctx, &fixtures.CoverLetterDraft)

	updated := fixtures.CoverLetterDraft
	updated.EditedText = "Edited content here"
	updated.Status = "approved"
	err := repo.Update(ctx, &updated)
	require.NoError(t, err)

	got, err := repo.GetByID(ctx, fixtures.CoverLetterDraftID)
	require.NoError(t, err)
	assert.Equal(t, "Edited content here", got.EditedText)
	assert.Equal(t, "approved", got.Status)
}

func TestCLGetByID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupCLRepo(t)

	_, err := repo.GetByID(context.Background(), uuid.New())
	require.ErrorIs(t, err, coverletter.ErrNotFound)
}
