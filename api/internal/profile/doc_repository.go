package profile

import (
	"context"
	"fmt"

	"bitbucket.org/irenato/coverletter-hub/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DocRepository interface {
	Create(ctx context.Context, doc *model.CVDocument) error
	GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]model.CVDocument, error)
}

type PostgresDocRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresDocRepository(pool *pgxpool.Pool) *PostgresDocRepository {
	return &PostgresDocRepository{pool: pool}
}

func (r *PostgresDocRepository) Create(ctx context.Context, doc *model.CVDocument) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO cv_documents (id, cv_profile_id, original_file, file_type, extracted_text, uploaded_at)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		doc.ID, doc.CVProfileID, doc.OriginalFile, doc.FileType, doc.ExtractedText, doc.UploadedAt,
	)
	if err != nil {
		return fmt.Errorf("inserting cv document: %w", err)
	}
	return nil
}

func (r *PostgresDocRepository) GetByProfileID(ctx context.Context, profileID uuid.UUID) ([]model.CVDocument, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, cv_profile_id, file_type, extracted_text, uploaded_at
		 FROM cv_documents WHERE cv_profile_id = $1 ORDER BY uploaded_at DESC`, profileID,
	)
	if err != nil {
		return nil, fmt.Errorf("querying cv documents: %w", err)
	}
	defer rows.Close()

	var docs []model.CVDocument
	for rows.Next() {
		var d model.CVDocument
		if err := rows.Scan(&d.ID, &d.CVProfileID, &d.FileType, &d.ExtractedText, &d.UploadedAt); err != nil {
			return nil, fmt.Errorf("scanning cv document: %w", err)
		}
		docs = append(docs, d)
	}
	return docs, nil
}
