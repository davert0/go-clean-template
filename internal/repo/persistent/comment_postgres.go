package persistent

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const _defaultEntityCap = 64

// CommentsRepo -.
type CommentsRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *CommentsRepo {
	return &CommentsRepo{pg}
}

// GetComments -.
func (r *CommentsRepo) GetComments(ctx context.Context, e entity.Entity) ([]entity.Comment, error) {
	var entityRefID int64
	sql, args, err := r.Builder.
		Select("id").
		From("entity").
		Where(squirrel.Eq{
			"entity_id":   e.EntityID,
			"entity_type": e.EntityType,
		}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("CommentsRepo - GetComments - r.Builder: %w", err)
	}
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&entityRefID)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("CommentsRepo - GetComments - r.Pool.QueryRow: %w", err)
	}

	sql, args, err = r.Builder.
		Select("entity_ref_id, text, created_by, created_at").
		From("comment").
		Where(squirrel.Eq{
			"entity_ref_id": entityRefID,
		}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("CommentsRepo - GetComments - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("CommentsRepo - GetComments - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Comment, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Comment{}

		err = rows.Scan(&e.EntityRefID, &e.Text, &e.CreatedBy, &e.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("CommentsRepo - GetComments - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// CreateComment -.
func (r *CommentsRepo) CreateComment(ctx context.Context, c entity.Comment, e entity.Entity) (entity.Comment, error) {
	c.CreatedAt = time.Now()
	entityRefID, err := r.getOrCreateIntityRefID(ctx, e)
	if err != nil {
		return c, fmt.Errorf("CommentsRepo - CreateComment - r.getOrCreateIntityRefID: %w", err)
	}

	c.EntityRefID = entityRefID
	sql, args, err := r.Builder.
		Insert("comment").
		Columns("text, created_by, created_at, entity_ref_id").
		Values(c.Text, c.CreatedBy, c.CreatedAt, c.EntityRefID).
		ToSql()
	if err != nil {
		return c, fmt.Errorf("CommentsRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return c, fmt.Errorf("CommentsRepo - Store - r.Pool.Exec: %w", err)
	}

	return c, nil
}

func (r *CommentsRepo) getOrCreateIntityRefID(ctx context.Context, e entity.Entity) (int64, error) {
	var entityRefID int64
	sql, args, err := r.Builder.
		Select("id").
		From("entity").
		Where(squirrel.Eq{
			"entity_id":   e.EntityID,
			"entity_type": e.EntityType,
		}).
		ToSql()
	if err != nil {
		return -1, fmt.Errorf("CommentsRepo - DoComment - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&entityRefID)
	switch err {
	case pgx.ErrNoRows:
		sql, args, err := r.Builder.
			Insert("entity").
			Columns("entity_id", "entity_type").
			Values(e.EntityID, e.EntityType).
			Suffix("RETURNING id").
			ToSql()
		if err != nil {
			return -1, fmt.Errorf("CommentsRepo - DoComment - r.Builder: %w", err)
		}
		err = r.Pool.QueryRow(ctx, sql, args...).Scan(&entityRefID)
		if err == nil {
			return entityRefID, nil
		}
		return -1, err
	case nil:
		return entityRefID, nil
	default:
		return -1, err
	}
}
