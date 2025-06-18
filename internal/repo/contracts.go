// Package repo implements application outer layer logic. Each logic group in own file.
package repo

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
)

type (
	// TranslationRepo -.
	CommentsRepo interface {
		Store(context.Context, entity.Comment) error
		GetComments(context.Context) ([]entity.Comment, error)
		DoComment(context.Context, entity.Comment, entity.Entity) (entity.Comment, error)
	}
)
