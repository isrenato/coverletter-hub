package database_test

import (
	"context"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/database"
	"bitbucket.org/irenato/coverletter-hub/api/testutil"
	"github.com/stretchr/testify/require"
)

func TestMigrate_AppliesAllMigrations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	tdb := testutil.NewTestDB(t)
	ctx := context.Background()

	err := database.Migrate(ctx, tdb.Pool)
	require.NoError(t, err)

	var tableCount int
	err = tdb.Pool.QueryRow(ctx,
		`SELECT count(*) FROM information_schema.tables
		 WHERE table_schema = 'public' AND table_type = 'BASE TABLE'
		 AND table_name != 'schema_migrations'`).Scan(&tableCount)
	require.NoError(t, err)
	require.Equal(t, 5, tableCount, "expected 5 tables after migration")
}

func TestMigrate_Idempotent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	tdb := testutil.NewTestDB(t)
	ctx := context.Background()

	err := database.Migrate(ctx, tdb.Pool)
	require.NoError(t, err)

	err = database.Migrate(ctx, tdb.Pool)
	require.NoError(t, err, "second migration run should be idempotent")
}
