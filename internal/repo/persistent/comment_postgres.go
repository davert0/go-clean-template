package persistent

import (
	"context"
	"fmt"
	"time"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/pkg/postgres"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type TranslationRepo struct {
	*postgres.Postgres
}

// New -.
func NewT(pg *postgres.Postgres) *TranslationRepo {
	return &TranslationRepo{pg}
}

// GetHistory -.
func (r *TranslationRepo) GetHistory(ctx context.Context) ([]entity.Translation, error) {
	sql, _, err := r.Builder.
		Select("source, destination, original, translation").
		From("history").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Translation, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Translation{}

		err = rows.Scan(&e.Source, &e.Destination, &e.Original, &e.Translation)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// Store -.
func (r *TranslationRepo) Store(ctx context.Context, t entity.Translation) error {
	sql, args, err := r.Builder.
		Insert("history").
		Columns("source, destination, original, translation").
		Values(t.Source, t.Destination, t.Original, t.Translation).
		ToSql()
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}

// TranslationRepo -.
type CommentsRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *CommentsRepo {
	return &CommentsRepo{pg}
}

// GetHistory -.
func (r *CommentsRepo) GetComments(ctx context.Context) ([]entity.Comment, error) {
	sql, _, err := r.Builder.
		Select("text, created_by, created_at").
		From("comment").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("CommentsRepo - GetHistory - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("CommentsRepo - GetHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Comment, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Comment{}

		err = rows.Scan(&e.Text, &e.CreatedBy, &e.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("CommentsRepo - GetHistory - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// Store -.
func (r *CommentsRepo) Store(ctx context.Context, c entity.Comment) error {
	sql, args, err := r.Builder.
		Insert("comment").
		Columns("text, created_by, created_at"). //, entity_ref_id").
		Values(c.Text, c.CreatedBy, c.CreatedAt /*, len(c.EntityID)+len(c.EntityType)*/).
		ToSql()
	if err != nil {
		return fmt.Errorf("CommentsRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CommentsRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *CommentsRepo) DoComment(ctx context.Context, c entity.Comment) (entity.Comment, error) {
	c.CreatedAt = time.Now()

	return c, nil
}
