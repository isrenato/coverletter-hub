package vacancy

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("vacancy not found")

type ListOptions struct {
	Limit  int
	Offset int
}

type Repository interface {
	Create(ctx context.Context, v *model.Vacancy) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Vacancy, error)
	List(ctx context.Context, userID uuid.UUID, opts ListOptions) ([]model.Vacancy, int, error)
	Update(ctx context.Context, v *model.Vacancy) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) Create(ctx context.Context, v *model.Vacancy) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO vacancies (id, user_id, title, company, description, location, linkedin_url, source, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		v.ID, v.UserID, v.Title, v.Company, v.Description, v.Location, v.LinkedInURL, v.Source, v.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("inserting vacancy: %w", err)
	}
	return nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Vacancy, error) {
	var v model.Vacancy
	err := r.pool.QueryRow(ctx,
		`SELECT id, user_id, title, company, description, location, linkedin_url, source, created_at
		 FROM vacancies WHERE id = $1`, id,
	).Scan(&v.ID, &v.UserID, &v.Title, &v.Company, &v.Description, &v.Location, &v.LinkedInURL, &v.Source, &v.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("querying vacancy: %w", err)
	}
	return &v, nil
}

func (r *PostgresRepository) List(ctx context.Context, userID uuid.UUID, opts ListOptions) ([]model.Vacancy, int, error) {
	var total int
	err := r.pool.QueryRow(ctx, "SELECT count(*) FROM vacancies WHERE user_id = $1", userID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("counting vacancies: %w", err)
	}

	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, title, company, description, location, linkedin_url, source, created_at
		 FROM vacancies WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
		userID, opts.Limit, opts.Offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("listing vacancies: %w", err)
	}
	defer rows.Close()

	var items []model.Vacancy
	for rows.Next() {
		var v model.Vacancy
		if err := rows.Scan(&v.ID, &v.UserID, &v.Title, &v.Company, &v.Description, &v.Location, &v.LinkedInURL, &v.Source, &v.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("scanning vacancy: %w", err)
		}
		items = append(items, v)
	}
	return items, total, nil
}

func (r *PostgresRepository) Update(ctx context.Context, v *model.Vacancy) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE vacancies SET title=$1, company=$2, description=$3, location=$4, linkedin_url=$5
		 WHERE id=$6`,
		v.Title, v.Company, v.Description, v.Location, v.LinkedInURL, v.ID,
	)
	if err != nil {
		return fmt.Errorf("updating vacancy: %w", err)
	}
	return nil
}

func (r *PostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tag, err := r.pool.Exec(ctx, "DELETE FROM vacancies WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("deleting vacancy: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
