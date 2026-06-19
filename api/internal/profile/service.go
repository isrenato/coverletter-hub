package profile

import (
	"context"
	"fmt"
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
)

type CVParser interface {
	ParseCV(ctx context.Context, fileContent []byte, fileType string) (*model.CVProfile, error)
}

type Service struct {
	profiles Repository
	docs     DocRepository
	parser   CVParser
}

func NewService(profiles Repository, docs DocRepository, parser CVParser) *Service {
	return &Service{profiles: profiles, docs: docs, parser: parser}
}

func (s *Service) Get(ctx context.Context, userID uuid.UUID) (*model.CVProfile, error) {
	return s.profiles.GetByUserID(ctx, userID)
}

func (s *Service) Update(ctx context.Context, userID uuid.UUID, p *model.CVProfile) (*model.CVProfile, error) {
	existing, err := s.profiles.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("fetching profile to update: %w", err)
	}

	existing.FullName = p.FullName
	existing.Headline = p.Headline
	existing.Summary = p.Summary
	existing.Experience = p.Experience
	existing.Education = p.Education
	existing.Skills = p.Skills
	existing.Languages = p.Languages

	if err := s.profiles.Update(ctx, existing); err != nil {
		return nil, fmt.Errorf("updating profile: %w", err)
	}
	return existing, nil
}

func (s *Service) Upload(ctx context.Context, userID uuid.UUID, fileContent []byte, fileType string) (*model.CVProfile, error) {
	parsed, err := s.parser.ParseCV(ctx, fileContent, fileType)
	if err != nil {
		return nil, fmt.Errorf("parsing CV: %w", err)
	}

	existing, err := s.profiles.GetByUserID(ctx, userID)
	if err != nil {
		now := time.Now()
		existing = &model.CVProfile{
			ID:         uuid.New(),
			UserID:     userID,
			FullName:   parsed.FullName,
			Headline:   parsed.Headline,
			Summary:    parsed.Summary,
			Experience: parsed.Experience,
			Education:  parsed.Education,
			Skills:     parsed.Skills,
			Languages:  parsed.Languages,
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		if err := s.profiles.Create(ctx, existing); err != nil {
			return nil, fmt.Errorf("creating profile: %w", err)
		}
	} else {
		existing.FullName = parsed.FullName
		existing.Headline = parsed.Headline
		existing.Summary = parsed.Summary
		existing.Experience = parsed.Experience
		existing.Education = parsed.Education
		existing.Skills = parsed.Skills
		existing.Languages = parsed.Languages
		if err := s.profiles.Update(ctx, existing); err != nil {
			return nil, fmt.Errorf("updating profile from upload: %w", err)
		}
	}

	doc := &model.CVDocument{
		ID:            uuid.New(),
		CVProfileID:   existing.ID,
		OriginalFile:  fileContent,
		FileType:      fileType,
		ExtractedText: fmt.Sprintf("Parsed: %s - %s", parsed.FullName, parsed.Headline),
		UploadedAt:    time.Now(),
	}
	if err := s.docs.Create(ctx, doc); err != nil {
		return nil, fmt.Errorf("storing document: %w", err)
	}

	return existing, nil
}
