package parser_test

import (
	"context"
	"encoding/json"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockLLM struct {
	response string
	err      error
}

func (m *mockLLM) Generate(_ context.Context, _, _ string) (string, error) {
	return m.response, m.err
}

func TestParseCV_ExtractsProfile(t *testing.T) {
	profileJSON, _ := json.Marshal(map[string]interface{}{
		"full_name":  "John Doe",
		"headline":   "Senior Software Engineer",
		"summary":    "10 years experience",
		"experience": []map[string]string{{"title": "Engineer", "company": "TechCorp"}},
		"education":  []map[string]string{{"degree": "BSc CS", "school": "MIT"}},
		"skills":     []string{"Go", "TypeScript"},
		"languages":  []string{"English"},
	})

	llm := &mockLLM{response: string(profileJSON)}
	p := parser.New(llm)

	result, err := p.ParseCV(context.Background(), []byte("fake pdf content"), "pdf")
	require.NoError(t, err)
	assert.Equal(t, "John Doe", result.FullName)
	assert.Equal(t, "Senior Software Engineer", result.Headline)
	assert.Equal(t, "10 years experience", result.Summary)
}

func TestParseCV_InvalidJSON(t *testing.T) {
	llm := &mockLLM{response: "not json at all"}
	p := parser.New(llm)

	_, err := p.ParseCV(context.Background(), []byte("content"), "pdf")
	require.Error(t, err)
}
