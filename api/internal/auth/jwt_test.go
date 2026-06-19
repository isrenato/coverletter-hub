package auth_test

import (
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/auth"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSecret = "test-jwt-secret-key"

func TestGenerateToken_And_Validate(t *testing.T) {
	token, err := auth.GenerateToken(fixtures.UserJohn, testSecret)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := auth.ValidateToken(token, testSecret)
	require.NoError(t, err)
	assert.Equal(t, fixtures.UserJohnID.String(), claims.UserID)
	assert.Equal(t, fixtures.UserJohn.Email, claims.Email)
}

func TestValidateToken_InvalidToken(t *testing.T) {
	_, err := auth.ValidateToken("invalid.token.here", testSecret)
	require.Error(t, err)
}

func TestValidateToken_WrongSecret(t *testing.T) {
	token, err := auth.GenerateToken(fixtures.UserJohn, testSecret)
	require.NoError(t, err)

	_, err = auth.ValidateToken(token, "wrong-secret")
	require.Error(t, err)
}
