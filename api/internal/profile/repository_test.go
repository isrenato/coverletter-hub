package profile_test

import (
	"context"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/profile"
	"bitbucket.org/irenato/coverletter-hub/api/testutil"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupProfileRepo(t *testing.T) (*profile.PostgresRepository, *testutil.TestDB) {
	tdb := testutil.NewTestDB(t)
	testutil.MigrateTestDB(t, tdb.Pool)
	testutil.SeedUser(t, tdb.Pool, fixtures.UserJohn)
	return profile.NewPostgresRepository(tdb.Pool), tdb
}

func TestProfileCreate_And_GetByUserID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupProfileRepo(t)
	ctx := context.Background()

	err := repo.Create(ctx, &fixtures.CVProfileJohn)
	require.NoError(t, err)

	got, err := repo.GetByUserID(ctx, fixtures.UserJohnID)
	require.NoError(t, err)
	assert.Equal(t, fixtures.CVProfileJohn.FullName, got.FullName)
	assert.Equal(t, fixtures.CVProfileJohn.Headline, got.Headline)
	assert.Equal(t, fixtures.CVProfileJohn.Summary, got.Summary)
}

func TestProfileGetByUserID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupProfileRepo(t)
	ctx := context.Background()

	_, err := repo.GetByUserID(ctx, fixtures.UserJaneID)
	require.ErrorIs(t, err, profile.ErrNotFound)
}

func TestProfileUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupProfileRepo(t)
	ctx := context.Background()

	err := repo.Create(ctx, &fixtures.CVProfileJohn)
	require.NoError(t, err)

	updated := fixtures.CVProfileJohn
	updated.Headline = "Staff Engineer"
	err = repo.Update(ctx, &updated)
	require.NoError(t, err)

	got, err := repo.GetByUserID(ctx, fixtures.UserJohnID)
	require.NoError(t, err)
	assert.Equal(t, "Staff Engineer", got.Headline)
}
