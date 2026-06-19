package user

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	Create(ctx context.Context, u *model.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByLinkedInID(ctx context.Context, linkedInID string) (*model.User, error)
	Update(ctx context.Context, u *model.User) error
}

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) Create(ctx context.Context, u *model.User) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO users (id, linkedin_id, email, name, access_token, refresh_token, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		u.ID, u.LinkedInID, u.Email, u.Name, u.AccessToken, u.RefreshToken, u.CreatedAt, u.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("inserting user: %w", err)
	}
	return nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var u model.User
	err := r.pool.QueryRow(ctx,
		`SELECT id, linkedin_id, email, name, access_token, refresh_token, created_at, updated_at
		 FROM users WHERE id = $1`, id,
	).Scan(&u.ID, &u.LinkedInID, &u.Email, &u.Name, &u.AccessToken, &u.RefreshToken, &u.CreatedAt, &u.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("querying user by id: %w", err)
	}
	return &u, nil
}

func (r *PostgresRepository) GetByLinkedInID(ctx context.Context, linkedInID string) (*model.User, error) {
	var u model.User
	err := r.pool.QueryRow(ctx,
		`SELECT id, linkedin_id, email, name, access_token, refresh_token, created_at, updated_at
		 FROM users WHERE linkedin_id = $1`, linkedInID,
	).Scan(&u.ID, &u.LinkedInID, &u.Email, &u.Name, &u.AccessToken, &u.RefreshToken, &u.CreatedAt, &u.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("querying user by linkedin id: %w", err)
	}
	return &u, nil
}

func (r *PostgresRepository) Update(ctx context.Context, u *model.User) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE users SET email = $1, name = $2, access_token = $3, refresh_token = $4, updated_at = NOW()
		 WHERE id = $5`,
		u.Email, u.Name, u.AccessToken, u.RefreshToken, u.ID,
	)
	if err != nil {
		return fmt.Errorf("updating user: %w", err)
	}
	return nil
}
