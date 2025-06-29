package v1

import (
	"github.com/evrone/go-clean-template/internal/usecase/comment/create"
	"github.com/evrone/go-clean-template/internal/usecase/comment/get"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// NewCommentRoutes -.
func NewCommentRoutes(apiV1Group fiber.Router, c create.UseCase, g get.UseCase, l logger.Interface) {
	router := &V1{createUC: c, getUC: g, logger: l, validator: validator.New(validator.WithRequiredStructEnabled())}

	commentGroup := apiV1Group.Group("/comment")

	{
		commentGroup.Get("/comments", router.getComments)
		commentGroup.Post("/create-comment", router.createComment)
	}
}
