package llm_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/llm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClaudeClient_Generate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Contains(t, r.Header.Get("x-api-key"), "test-key")
		assert.Equal(t, "2023-06-01", r.Header.Get("anthropic-version"))

		var body map[string]interface{}
		json.NewDecoder(r.Body).Decode(&body)
		assert.Equal(t, "claude-sonnet-4-20250514", body["model"])

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"content": []map[string]interface{}{
				{"type": "text", "text": "Generated response text"},
			},
		})
	}))
	defer server.Close()

	client := llm.NewClaudeClient("test-key")
	client.BaseURL = server.URL

	result, err := client.Generate(context.Background(), "You are a helpful assistant", "Write a cover letter")
	require.NoError(t, err)
	assert.Equal(t, "Generated response text", result)
}

func TestClaudeClient_Generate_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"error":{"message":"rate limited"}}`))
	}))
	defer server.Close()

	client := llm.NewClaudeClient("test-key")
	client.BaseURL = server.URL

	_, err := client.Generate(context.Background(), "system", "user")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "429")
}
