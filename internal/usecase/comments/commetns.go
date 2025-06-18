package comments

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/internal/repo"
)

// UseCase -.
type UseCase struct {
	repo repo.CommentsRepo
}

// New -.
func New(r repo.CommentsRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

// History - getting comments history from store.
func (uc *UseCase) History(ctx context.Context) (entity.CommentsHistory, error) {
	comments, err := uc.repo.GetComments(ctx)
	if err != nil {
		return entity.CommentsHistory{}, fmt.Errorf("CommentsUseCase - History - s.repo.GetComments: %w", err)
	}

	return entity.CommentsHistory{Comments: comments}, nil
}

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
