package coverletter

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("cover letter not found")

type ListOptions struct {
	Limit  int
	Offset int
	Status string
}

type Repository interface {
	Create(ctx context.Context, cl *model.CoverLetter) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.CoverLetter, error)
	List(ctx context.Context, userID uuid.UUID, opts ListOptions) ([]model.CoverLetter, int, error)
	Update(ctx context.Context, cl *model.CoverLetter) error
	GetByVacancyID(ctx context.Context, vacancyID uuid.UUID) ([]model.CoverLetter, error)
}

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) Create(ctx context.Context, cl *model.CoverLetter) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO cover_letters (id, vacancy_id, cv_profile_id, generated_text, edited_text, status, generated_at, approved_at, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		cl.ID, cl.VacancyID, cl.CVProfileID, cl.GeneratedText, cl.EditedText, cl.Status, cl.GeneratedAt, cl.ApprovedAt, cl.CreatedAt, cl.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("inserting cover letter: %w", err)
	}
	return nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.CoverLetter, error) {
	var cl model.CoverLetter
	err := r.pool.QueryRow(ctx,
		`SELECT id, vacancy_id, cv_profile_id, generated_text, edited_text, status, generated_at, approved_at, created_at, updated_at
		 FROM cover_letters WHERE id = $1`, id,
	).Scan(&cl.ID, &cl.VacancyID, &cl.CVProfileID, &cl.GeneratedText, &cl.EditedText, &cl.Status, &cl.GeneratedAt, &cl.ApprovedAt, &cl.CreatedAt, &cl.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("querying cover letter: %w", err)
	}
	return &cl, nil
}

func (r *PostgresRepository) List(ctx context.Context, userID uuid.UUID, opts ListOptions) ([]model.CoverLetter, int, error) {
	baseWhere := `FROM cover_letters cl JOIN vacancies v ON cl.vacancy_id = v.id WHERE v.user_id = $1`
	args := []interface{}{userID}

	if opts.Status != "" {
		baseWhere += ` AND cl.status = $2`
		args = append(args, opts.Status)
	}

	var total int
	err := r.pool.QueryRow(ctx, "SELECT count(*) "+baseWhere, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("counting cover letters: %w", err)
	}

	limitIdx := len(args) + 1
	offsetIdx := len(args) + 2
	query := fmt.Sprintf(
		`SELECT cl.id, cl.vacancy_id, cl.cv_profile_id, cl.generated_text, cl.edited_text, cl.status, cl.generated_at, cl.approved_at, cl.created_at, cl.updated_at
		 %s ORDER BY cl.created_at DESC LIMIT $%d OFFSET $%d`, baseWhere, limitIdx, offsetIdx)
	args = append(args, opts.Limit, opts.Offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("listing cover letters: %w", err)
	}
	defer rows.Close()

	var items []model.CoverLetter
	for rows.Next() {
		var cl model.CoverLetter
		if err := rows.Scan(&cl.ID, &cl.VacancyID, &cl.CVProfileID, &cl.GeneratedText, &cl.EditedText, &cl.Status, &cl.GeneratedAt, &cl.ApprovedAt, &cl.CreatedAt, &cl.UpdatedAt); err != nil {
			return nil, 0, fmt.Errorf("scanning cover letter: %w", err)
		}
		items = append(items, cl)
	}
	return items, total, nil
}

func (r *PostgresRepository) Update(ctx context.Context, cl *model.CoverLetter) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE cover_letters SET edited_text=$1, status=$2, approved_at=$3, updated_at=NOW()
		 WHERE id=$4`,
		cl.EditedText, cl.Status, cl.ApprovedAt, cl.ID,
	)
	if err != nil {
		return fmt.Errorf("updating cover letter: %w", err)
	}
	return nil
}

func (r *PostgresRepository) GetByVacancyID(ctx context.Context, vacancyID uuid.UUID) ([]model.CoverLetter, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, vacancy_id, cv_profile_id, generated_text, edited_text, status, generated_at, approved_at, created_at, updated_at
		 FROM cover_letters WHERE vacancy_id = $1 ORDER BY created_at DESC`, vacancyID,
	)
	if err != nil {
		return nil, fmt.Errorf("querying cover letters by vacancy: %w", err)
	}
	defer rows.Close()

	var items []model.CoverLetter
	for rows.Next() {
		var cl model.CoverLetter
		if err := rows.Scan(&cl.ID, &cl.VacancyID, &cl.CVProfileID, &cl.GeneratedText, &cl.EditedText, &cl.Status, &cl.GeneratedAt, &cl.ApprovedAt, &cl.CreatedAt, &cl.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scanning cover letter: %w", err)
		}
		items = append(items, cl)
	}
	return items, nil
}
