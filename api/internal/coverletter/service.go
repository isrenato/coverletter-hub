package coverletter

import (
	"context"
	"fmt"
	"time"

	"bitbucket.org/irenato/coverletter-hub/api/internal/llm"
	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
)

type VacancyGetter interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.Vacancy, error)
}

type ProfileGetter interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*model.CVProfile, error)
}

type GenerateResult struct {
	model.CoverLetter
	HasWarning bool   `json:"has_warning,omitempty"`
	Warning    string `json:"warning,omitempty"`
}

type Service struct {
	repo      Repository
	vacancies VacancyGetter
	profiles  ProfileGetter
	llm       llm.Client
}

func NewService(repo Repository, vacancies VacancyGetter, profiles ProfileGetter, llmClient llm.Client) *Service {
	return &Service{repo: repo, vacancies: vacancies, profiles: profiles, llm: llmClient}
}

func (s *Service) Generate(ctx context.Context, userID, vacancyID uuid.UUID) (*GenerateResult, error) {
	vacancy, err := s.vacancies.GetByID(ctx, vacancyID)
	if err != nil {
		return nil, fmt.Errorf("fetching vacancy: %w", err)
	}

	profile, err := s.profiles.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("fetching profile: %w", err)
	}

	result := &GenerateResult{}

	existing, err := s.repo.GetByVacancyID(ctx, vacancyID)
	if err == nil {
		sixMonthsAgo := time.Now().AddDate(0, -6, 0)
		for _, cl := range existing {
			if cl.Status == model.CoverLetterStatusApproved && cl.ApprovedAt != nil && cl.ApprovedAt.After(sixMonthsAgo) {
				result.HasWarning = true
				result.Warning = "You have an approved cover letter for this vacancy from less than 6 months ago"
				break
			}
		}
	}

	prompt := buildUserPrompt(profile, vacancy)
	generated, err := s.llm.Generate(ctx, coverLetterSystemPrompt, prompt)
	if err != nil {
		return nil, fmt.Errorf("generating cover letter: %w", err)
	}

	now := time.Now()
	cl := model.CoverLetter{
		ID:            uuid.New(),
		VacancyID:     vacancyID,
		CVProfileID:   profile.ID,
		GeneratedText: generated,
		Status:        model.CoverLetterStatusDraft,
		GeneratedAt:   now,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := s.repo.Create(ctx, &cl); err != nil {
		return nil, fmt.Errorf("saving cover letter: %w", err)
	}

	result.CoverLetter = cl
	return result, nil
}

func (s *Service) UpdateText(ctx context.Context, id uuid.UUID, editedText string) (*model.CoverLetter, error) {
	cl, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetching cover letter: %w", err)
	}

	cl.EditedText = editedText
	if err := s.repo.Update(ctx, cl); err != nil {
		return nil, fmt.Errorf("updating cover letter text: %w", err)
	}
	return cl, nil
}

func (s *Service) UpdateStatus(ctx context.Context, id uuid.UUID, status string) (*model.CoverLetter, error) {
	cl, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("fetching cover letter: %w", err)
	}

	cl.Status = status
	if status == model.CoverLetterStatusApproved {
		now := time.Now()
		cl.ApprovedAt = &now
	}

	if err := s.repo.Update(ctx, cl); err != nil {
		return nil, fmt.Errorf("updating cover letter status: %w", err)
	}
	return cl, nil
}
