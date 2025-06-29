package v1

import (
	"github.com/evrone/go-clean-template/internal/usecase/comment/create"
	"github.com/evrone/go-clean-template/internal/usecase/comment/get"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	createUC  create.UseCase
	getUC     get.UseCase
	logger    logger.Interface
	validator *validator.Validate
}
