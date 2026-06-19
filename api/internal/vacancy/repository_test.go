package vacancy_test

import (
	"context"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/vacancy"
	"bitbucket.org/irenato/coverletter-hub/api/testutil"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupVacancyRepo(t *testing.T) (*vacancy.PostgresRepository, *testutil.TestDB) {
	tdb := testutil.NewTestDB(t)
	testutil.MigrateTestDB(t, tdb.Pool)
	testutil.SeedUser(t, tdb.Pool, fixtures.UserJohn)
	return vacancy.NewPostgresRepository(tdb.Pool), tdb
}

func TestVacancyCreate_And_GetByID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupVacancyRepo(t)
	ctx := context.Background()

	err := repo.Create(ctx, &fixtures.VacancyBackend)
	require.NoError(t, err)

	got, err := repo.GetByID(ctx, fixtures.VacancyBackendID)
	require.NoError(t, err)
	assert.Equal(t, "Backend Engineer", got.Title)
	assert.Equal(t, "StartupCo", got.Company)
}

func TestVacancyList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupVacancyRepo(t)
	ctx := context.Background()

	_ = repo.Create(ctx, &fixtures.VacancyBackend)
	_ = repo.Create(ctx, &fixtures.VacancyFrontend)

	items, total, err := repo.List(ctx, fixtures.UserJohnID, vacancy.ListOptions{Limit: 10, Offset: 0})
	require.NoError(t, err)
	assert.Equal(t, 2, total)
	assert.Len(t, items, 2)
}

func TestVacancyDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupVacancyRepo(t)
	ctx := context.Background()

	_ = repo.Create(ctx, &fixtures.VacancyBackend)
	err := repo.Delete(ctx, fixtures.VacancyBackendID)
	require.NoError(t, err)

	_, err = repo.GetByID(ctx, fixtures.VacancyBackendID)
	require.ErrorIs(t, err, vacancy.ErrNotFound)
}

func TestVacancyGetByID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupVacancyRepo(t)
	ctx := context.Background()

	_, err := repo.GetByID(ctx, uuid.New())
	require.ErrorIs(t, err, vacancy.ErrNotFound)
}
