<<<<<<<< HEAD:internal/usecase/comment/get/usecase.go
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
========
package create

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
>>>>>>>> remotes/origin/comment_handle:internal/usecase/comment/create/usecase.go
