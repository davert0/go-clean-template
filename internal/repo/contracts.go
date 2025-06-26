// Package repo implements application outer layer logic. Each logic group in own file.
package repo

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
)

type (
	// CommentsRepo -.
	CommentsRepo interface {
		CreateComment(context.Context, entity.Comment, entity.Entity) (entity.Comment, error)
		GetComments(context.Context, entity.Entity, string) ([]entity.Comment, error)
	}
)
