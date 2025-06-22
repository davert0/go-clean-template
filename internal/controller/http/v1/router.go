package v1

import (
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(apiV1Group fiber.Router, c usecase.Comment, l logger.Interface) {
	router := &V1{comment: c, logger: l, validator: validator.New(validator.WithRequiredStructEnabled())}

	commentGroup := apiV1Group.Group("/comment")

	{
		commentGroup.Get("/comments", router.getComments)
		commentGroup.Post("/do-comment", router.doComment)
	}
}
