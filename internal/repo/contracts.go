// Package repo implements application outer layer logic. Each logic group in own file.
package repo

import (
	"context"

	"github.com/evrone/go-clean-template/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_repo_test.go -package=usecase_test

type (
	// TranslationRepo -.
	TranslationRepo interface {
		Store(context.Context, entity.Translation) error
		GetHistory(context.Context) ([]entity.Translation, error)
	}

	// TranslationWebAPI -.
	TranslationWebAPI interface {
		Translate(entity.Translation) (entity.Translation, error)
	}
)

type (
	// TranslationRepo -.
	CommentsRepo interface {
		Store(context.Context, entity.Comment) error
		GetComments(context.Context) ([]entity.Comment, error)
		DoComment(context.Context, entity.Comment, entity.Entity) (entity.Comment, error)
	}
)
