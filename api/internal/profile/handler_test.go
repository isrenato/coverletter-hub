package profile_test

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/auth"
	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"bitbucket.org/irenato/coverletter-hub/api/internal/profile"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleGet_ReturnsProfile(t *testing.T) {
	repo := newMockProfileRepo()
	_ = repo.Create(context.Background(), &fixtures.CVProfileJohn)

	svc := profile.NewService(repo, &mockDocRepo{}, &mockParser{})
	h := profile.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/api/profile", nil)
	ctx := auth.WithUserID(req.Context(), fixtures.UserJohnID.String())
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	h.HandleGet(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var got model.CVProfile
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))
	assert.Equal(t, "John Doe", got.FullName)
}

func TestHandleGet_Unauthorized(t *testing.T) {
	svc := profile.NewService(newMockProfileRepo(), &mockDocRepo{}, &mockParser{})
	h := profile.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/api/profile", nil)
	rr := httptest.NewRecorder()
	h.HandleGet(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestHandleUpload_CreatesProfile(t *testing.T) {
	parsed := &model.CVProfile{
		FullName:   "Uploaded User",
		Headline:   "Engineer",
		Summary:    "Summary",
		Experience: json.RawMessage(`[]`),
		Education:  json.RawMessage(`[]`),
		Skills:     json.RawMessage(`[]`),
		Languages:  json.RawMessage(`[]`),
	}
	svc := profile.NewService(newMockProfileRepo(), &mockDocRepo{}, &mockParser{result: parsed})
	h := profile.NewHandler(svc)

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("cv", "resume.pdf")
	require.NoError(t, err)
	_, err = part.Write([]byte("fake pdf content"))
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	req := httptest.NewRequest(http.MethodPost, "/api/profile/upload", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	ctx := auth.WithUserID(req.Context(), fixtures.UserJohnID.String())
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	h.HandleUpload(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var got model.CVProfile
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))
	assert.Equal(t, "Uploaded User", got.FullName)
}
