package auth_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinkedInClient_AuthURL(t *testing.T) {
	client := auth.NewLinkedInClient("client-id", "client-secret", "http://localhost/callback")
	url := client.AuthURL()
	assert.Contains(t, url, "linkedin.com")
	assert.Contains(t, url, "client_id=client-id")
	assert.Contains(t, url, "redirect_uri=http")
}

func TestLinkedInClient_ExchangeCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token":  "test-access-token",
			"expires_in":    3600,
			"refresh_token": "test-refresh-token",
		})
	}))
	defer server.Close()

	client := auth.NewLinkedInClient("client-id", "client-secret", "http://localhost/callback")
	client.TokenURL = server.URL

	resp, err := client.ExchangeCode(context.Background(), "test-code")
	require.NoError(t, err)
	assert.Equal(t, "test-access-token", resp.AccessToken)
	assert.Equal(t, "test-refresh-token", resp.RefreshToken)
}

func TestLinkedInClient_GetProfile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
		json.NewEncoder(w).Encode(map[string]interface{}{
			"sub":   "linkedin-id-123",
			"name":  "Test User",
			"email": "test@example.com",
		})
	}))
	defer server.Close()

	client := auth.NewLinkedInClient("client-id", "client-secret", "http://localhost/callback")
	client.ProfileURL = server.URL

	profile, err := client.GetProfile(context.Background(), "test-token")
	require.NoError(t, err)
	assert.Equal(t, "linkedin-id-123", profile.Sub)
	assert.Equal(t, "Test User", profile.Name)
	assert.Equal(t, "test@example.com", profile.Email)
}
