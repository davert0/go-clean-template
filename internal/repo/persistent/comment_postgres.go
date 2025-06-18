package persistent

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/pkg/postgres"
)

const _defaultEntityCap = 64

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
		Select("entity_ref_id, text, created_by, created_at").
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

		err = rows.Scan(&e.EntityRefID, &e.Text, &e.CreatedBy, &e.CreatedAt)
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
		Columns("text, created_by, created_at, entity_ref_id").
		Values(c.Text, c.CreatedBy, c.CreatedAt, c.EntityRefID).
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

func (r *CommentsRepo) DoComment(ctx context.Context, c entity.Comment, e entity.Entity) (entity.Comment, error) {
	c.CreatedAt = time.Now()

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
		return c, fmt.Errorf("CommentsRepo - DoComment - r.Builder: %w", err)
	}
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&entityRefID)

	if err == nil {
		c.EntityRefID = entityRefID
		return c, nil
	} else {
		sql, args, err := r.Builder.
			Insert("entity").
			Columns("entity_id", "entity_type").
			Values(e.EntityID, e.EntityType).
			Suffix("RETURNING id").
			ToSql()
		if err != nil {
			return c, fmt.Errorf("CommentsRepo - DoComment - r.Builder: %w", err)
		}
		err = r.Pool.QueryRow(ctx, sql, args...).Scan(&entityRefID)
		if err == nil {
			c.EntityRefID = entityRefID
			return c, nil
		}
		return c, err
	}
}
