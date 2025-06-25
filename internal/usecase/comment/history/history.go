package history

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
)

// History - getting all comments history from store fore entity.
func (uc *UseCase) GetComments(ctx context.Context, e entity.Entity) (entity.CommentsList, error) {
	comments, err := uc.repo.GetComments(ctx, e)
	if err != nil {
		return entity.CommentsList{}, fmt.Errorf("CommentsUseCase - History - s.repo.GetComments: %w", err)
	}

	return entity.CommentsList{EntityID: e.EntityID, EntityType: e.EntityType, Comments: comments}, nil
}
