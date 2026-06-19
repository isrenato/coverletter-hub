package testutil

import (
	"context"
	"testing"

	"bitbucket.org/irenato/coverletter-hub/api/internal/database"
	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MigrateTestDB(t *testing.T, pool *pgxpool.Pool) {
	t.Helper()
	ctx := context.Background()
	if err := database.Migrate(ctx, pool); err != nil {
		t.Fatalf("failed to run migrations: %v", err)
	}
}

func SeedUser(t *testing.T, pool *pgxpool.Pool, u model.User) {
	t.Helper()
	ctx := context.Background()
	_, err := pool.Exec(ctx,
		`INSERT INTO users (id, linkedin_id, email, name, access_token, refresh_token, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		u.ID, u.LinkedInID, u.Email, u.Name, u.AccessToken, u.RefreshToken, u.CreatedAt, u.UpdatedAt,
	)
	if err != nil {
		t.Fatalf("failed to seed user: %v", err)
	}
}

func SeedCVProfile(t *testing.T, pool *pgxpool.Pool, p model.CVProfile) {
	t.Helper()
	ctx := context.Background()
	_, err := pool.Exec(ctx,
		`INSERT INTO cv_profiles (id, user_id, full_name, headline, summary, experience, education, skills, languages, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		p.ID, p.UserID, p.FullName, p.Headline, p.Summary, p.Experience, p.Education, p.Skills, p.Languages, p.CreatedAt, p.UpdatedAt,
	)
	if err != nil {
		t.Fatalf("failed to seed cv profile: %v", err)
	}
}

func SeedVacancy(t *testing.T, pool *pgxpool.Pool, v model.Vacancy) {
	t.Helper()
	ctx := context.Background()
	_, err := pool.Exec(ctx,
		`INSERT INTO vacancies (id, user_id, title, company, description, location, linkedin_url, source, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		v.ID, v.UserID, v.Title, v.Company, v.Description, v.Location, v.LinkedInURL, v.Source, v.CreatedAt,
	)
	if err != nil {
		t.Fatalf("failed to seed vacancy: %v", err)
	}
}

func SeedCoverLetter(t *testing.T, pool *pgxpool.Pool, cl model.CoverLetter) {
	t.Helper()
	ctx := context.Background()
	_, err := pool.Exec(ctx,
		`INSERT INTO cover_letters (id, vacancy_id, cv_profile_id, generated_text, edited_text, status, generated_at, approved_at, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		cl.ID, cl.VacancyID, cl.CVProfileID, cl.GeneratedText, cl.EditedText, cl.Status, cl.GeneratedAt, cl.ApprovedAt, cl.CreatedAt, cl.UpdatedAt,
	)
	if err != nil {
		t.Fatalf("failed to seed cover letter: %v", err)
	}
}
