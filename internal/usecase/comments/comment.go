package comments

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
)

// Comment -.
func (uc *UseCase) Comment(ctx context.Context, c entity.Comment, e entity.Entity) (entity.Comment, error) {
	comment, err := uc.repo.DoComment(ctx, c, e)
	if err != nil {
		return entity.Comment{}, fmt.Errorf("CommentUseCase - Comment - DoComment: %w", err)
	}

	err = uc.repo.Store(ctx, comment)
	if err != nil {
		return entity.Comment{}, fmt.Errorf("CommentUseCase - Comment - s.repo.Store: %w", err)
	}

	return comment, nil
}
