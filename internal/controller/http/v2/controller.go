package v2

import (
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V2 -.
type V2 struct {
	t usecase.Comment
	l logger.Interface
	v *validator.Validate
}
