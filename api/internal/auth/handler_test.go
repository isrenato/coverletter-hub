package auth_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/auth"
	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"bitbucket.org/irenato/coverletter-hub/api/testutil/fixtures"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockUserRepo struct {
	users map[string]*model.User
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{users: make(map[string]*model.User)}
}

func (m *mockUserRepo) Create(_ context.Context, u *model.User) error {
	m.users[u.LinkedInID] = u
	return nil
}

func (m *mockUserRepo) GetByID(_ context.Context, id uuid.UUID) (*model.User, error) {
	for _, u := range m.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (m *mockUserRepo) GetByLinkedInID(_ context.Context, linkedInID string) (*model.User, error) {
	if u, ok := m.users[linkedInID]; ok {
		return u, nil
	}
	return nil, fmt.Errorf("not found")
}

func (m *mockUserRepo) Update(_ context.Context, u *model.User) error {
	m.users[u.LinkedInID] = u
	return nil
}

func TestHandleLinkedInRedirect(t *testing.T) {
	linkedIn := auth.NewLinkedInClient("test-client-id", "test-secret", "http://localhost/callback")
	handler := auth.NewAuthHandler(newMockUserRepo(), linkedIn, testSecret, "http://localhost:3000")

	req := httptest.NewRequest(http.MethodGet, "/api/v1/auth/linkedin", nil)
	w := httptest.NewRecorder()

	handler.HandleLinkedInRedirect(w, req)

	assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
	location := w.Header().Get("Location")
	assert.Contains(t, location, "linkedin.com")
	assert.Contains(t, location, "client_id=test-client-id")
}

func TestHandleMe(t *testing.T) {
	repo := newMockUserRepo()
	u := fixtures.UserJohn
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	_ = repo.Create(context.Background(), &u)

	linkedIn := auth.NewLinkedInClient("id", "secret", "http://localhost/callback")
	handler := auth.NewAuthHandler(repo, linkedIn, testSecret, "http://localhost:3000")

	token, _ := auth.GenerateToken(u, testSecret)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/auth/me", nil)
	req = req.WithContext(auth.WithUserID(req.Context(), u.ID.String()))
	w := httptest.NewRecorder()

	handler.HandleMe(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp model.User
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, u.Email, resp.Email)
	_ = token
}
