package user_test

import (
	"context"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/user"
	"bitbucket.org/irenato/coverletter-hub/api/testutil"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupRepo(t *testing.T) (*user.PostgresRepository, *testutil.TestDB) {
	tdb := testutil.NewTestDB(t)
	testutil.MigrateTestDB(t, tdb.Pool)
	repo := user.NewPostgresRepository(tdb.Pool)
	return repo, tdb
}

func TestCreate_And_GetByID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupRepo(t)
	ctx := context.Background()

	err := repo.Create(ctx, &fixtures.UserJohn)
	require.NoError(t, err)

	got, err := repo.GetByID(ctx, fixtures.UserJohnID)
	require.NoError(t, err)
	assert.Equal(t, fixtures.UserJohn.Email, got.Email)
	assert.Equal(t, fixtures.UserJohn.Name, got.Name)
	assert.Equal(t, fixtures.UserJohn.LinkedInID, got.LinkedInID)
}

func TestGetByLinkedInID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupRepo(t)
	ctx := context.Background()

	err := repo.Create(ctx, &fixtures.UserJohn)
	require.NoError(t, err)

	got, err := repo.GetByLinkedInID(ctx, fixtures.UserJohn.LinkedInID)
	require.NoError(t, err)
	assert.Equal(t, fixtures.UserJohnID, got.ID)
}

func TestGetByID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupRepo(t)
	ctx := context.Background()

	_, err := repo.GetByID(ctx, uuid.New())
	require.ErrorIs(t, err, user.ErrNotFound)
}

func TestUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupRepo(t)
	ctx := context.Background()

	err := repo.Create(ctx, &fixtures.UserJohn)
	require.NoError(t, err)

	updated := fixtures.UserJohn
	updated.Name = "John Updated"
	updated.Email = "john.updated@example.com"
	err = repo.Update(ctx, &updated)
	require.NoError(t, err)

	got, err := repo.GetByID(ctx, fixtures.UserJohnID)
	require.NoError(t, err)
	assert.Equal(t, "John Updated", got.Name)
	assert.Equal(t, "john.updated@example.com", got.Email)
}
