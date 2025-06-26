package get

import (
	"context"
	"fmt"

	"github.com/evrone/go-clean-template/internal/entity"
)

// GetComments - getting all comments history from store fore entity.
func (uc *UseCase) GetComments(ctx context.Context, e entity.Entity, limit int, offset int, orderBy string) (entity.CommentsList, error) {
	comments, err := uc.repo.GetComments(ctx, e, limit, offset, orderBy)
	if err != nil {
		return entity.CommentsList{}, fmt.Errorf("CommentsUseCase - GetComments - uc.repo.GetComments: %w", err)
	}

	return entity.CommentsList{EntityID: e.EntityID, EntityType: e.EntityType, Comments: comments}, nil
}
