package profile

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("cv profile not found")

type Repository interface {
	Create(ctx context.Context, p *model.CVProfile) error
	GetByUserID(ctx context.Context, userID uuid.UUID) (*model.CVProfile, error)
	Update(ctx context.Context, p *model.CVProfile) error
}

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) Create(ctx context.Context, p *model.CVProfile) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO cv_profiles (id, user_id, full_name, headline, summary, experience, education, skills, languages, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		p.ID, p.UserID, p.FullName, p.Headline, p.Summary, p.Experience, p.Education, p.Skills, p.Languages, p.CreatedAt, p.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("inserting cv profile: %w", err)
	}
	return nil
}

func (r *PostgresRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*model.CVProfile, error) {
	var p model.CVProfile
	err := r.pool.QueryRow(ctx,
		`SELECT id, user_id, full_name, headline, summary, experience, education, skills, languages, created_at, updated_at
		 FROM cv_profiles WHERE user_id = $1`, userID,
	).Scan(&p.ID, &p.UserID, &p.FullName, &p.Headline, &p.Summary, &p.Experience, &p.Education, &p.Skills, &p.Languages, &p.CreatedAt, &p.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("querying cv profile: %w", err)
	}
	return &p, nil
}

func (r *PostgresRepository) Update(ctx context.Context, p *model.CVProfile) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE cv_profiles SET full_name=$1, headline=$2, summary=$3, experience=$4, education=$5, skills=$6, languages=$7, updated_at=NOW()
		 WHERE id=$8`,
		p.FullName, p.Headline, p.Summary, p.Experience, p.Education, p.Skills, p.Languages, p.ID,
	)
	if err != nil {
		return fmt.Errorf("updating cv profile: %w", err)
	}
	return nil
}
