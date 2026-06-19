package parser

import (
	"context"
	"encoding/json"
	"fmt"

	"bitbucket.org/irenato/coverletter-hub/api/internal/llm"
	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
)

type Parser struct {
	llm llm.Client
}

func New(client llm.Client) *Parser {
	return &Parser{llm: client}
}

type parsedProfile struct {
	FullName   string          `json:"full_name"`
	Headline   string          `json:"headline"`
	Summary    string          `json:"summary"`
	Experience json.RawMessage `json:"experience"`
	Education  json.RawMessage `json:"education"`
	Skills     json.RawMessage `json:"skills"`
	Languages  json.RawMessage `json:"languages"`
}

const systemPrompt = `You are a CV parser. Extract structured data from the provided CV text and return ONLY valid JSON with these fields:
- full_name (string)
- headline (string)
- summary (string)
- experience (array of objects with: title, company, start, end)
- education (array of objects with: degree, school, year)
- skills (array of strings)
- languages (array of strings)
Return ONLY the JSON object, no markdown or explanation.`

func (p *Parser) ParseCV(ctx context.Context, fileContent []byte, fileType string) (*model.CVProfile, error) {
	userPrompt := fmt.Sprintf("Parse this %s CV document:\n\n%s", fileType, string(fileContent))

	response, err := p.llm.Generate(ctx, systemPrompt, userPrompt)
	if err != nil {
		return nil, fmt.Errorf("calling LLM for CV parsing: %w", err)
	}

	var parsed parsedProfile
	if err := json.Unmarshal([]byte(response), &parsed); err != nil {
		return nil, fmt.Errorf("parsing LLM response as JSON: %w", err)
	}

	profile := &model.CVProfile{
		FullName:   parsed.FullName,
		Headline:   parsed.Headline,
		Summary:    parsed.Summary,
		Experience: parsed.Experience,
		Education:  parsed.Education,
		Skills:     parsed.Skills,
		Languages:  parsed.Languages,
	}

	return profile, nil
}
