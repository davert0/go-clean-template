package comments

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
)

// Comment -.
func (uc *UseCase) CreateComment(ctx context.Context, c entity.Comment, e entity.Entity) (entity.Comment, error) {
	comment, err := uc.repo.CreateComment(ctx, c, e)
	if err != nil {
		return entity.Comment{}, fmt.Errorf("CommentUseCase - Comment - uc.repo.CreateComment: %w", err)
	}

	return comment, nil
}
