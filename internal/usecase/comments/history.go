package comments

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
)

// History - getting all comments history from store.
func (uc *UseCase) History(ctx context.Context) (entity.CommentsList, error) {
	comments, err := uc.repo.GetComments(ctx)
	if err != nil {
		return entity.CommentsList{}, fmt.Errorf("CommentsUseCase - History - s.repo.GetComments: %w", err)
	}

	return entity.CommentsList{Comments: comments}, nil
}
