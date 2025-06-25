package v1

import (
	"github.com/evrone/go-clean-template/internal/usecase/comment/create"
	"github.com/evrone/go-clean-template/internal/usecase/comment/history"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// NewCommentRoutes -.
func NewCommentRoutes(apiV1Group fiber.Router, c create.UseCase, h history.UseCase, l logger.Interface) {
	router := &V1{createUC: c, getUC: h, logger: l, validator: validator.New(validator.WithRequiredStructEnabled())}

	commentGroup := apiV1Group.Group("/comment")

	{
		commentGroup.Get("/comments", router.getComments)
		commentGroup.Post("/do-comment", router.doComment)
	}
}
