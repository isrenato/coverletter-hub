package profile_test

import (
	"context"
	"testing"
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"bitbucket.org/irenato/coverletter-hub/api/internal/profile"
	"bitbucket.org/irenato/coverletter-hub/api/testutil"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupDocRepo(t *testing.T) (*profile.PostgresDocRepository, *testutil.TestDB) {
	tdb := testutil.NewTestDB(t)
	testutil.MigrateTestDB(t, tdb.Pool)
	testutil.SeedUser(t, tdb.Pool, fixtures.UserJohn)
	testutil.SeedCVProfile(t, tdb.Pool, fixtures.CVProfileJohn)
	return profile.NewPostgresDocRepository(tdb.Pool), tdb
}

func TestDocCreate_And_GetByProfileID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	repo, _ := setupDocRepo(t)
	ctx := context.Background()

	doc := model.CVDocument{
		ID:            uuid.New(),
		CVProfileID:   fixtures.CVProfileJohnID,
		OriginalFile:  []byte("fake pdf content"),
		FileType:      "pdf",
		ExtractedText: "John Doe - Senior Engineer",
		UploadedAt:    time.Now(),
	}

	err := repo.Create(ctx, &doc)
	require.NoError(t, err)

	docs, err := repo.GetByProfileID(ctx, fixtures.CVProfileJohnID)
	require.NoError(t, err)
	require.Len(t, docs, 1)
	assert.Equal(t, "pdf", docs[0].FileType)
	assert.Equal(t, "John Doe - Senior Engineer", docs[0].ExtractedText)
}
