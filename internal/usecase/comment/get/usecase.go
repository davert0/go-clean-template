package get

import (
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
