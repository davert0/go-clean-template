// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	Comment interface {
		CreateComment(context.Context, entity.Comment, entity.Entity) (entity.Comment, error)
		GetComments(context.Context, entity.Entity) (entity.CommentsList, error)
	}
)
