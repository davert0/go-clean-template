package v1

import (
	"fmt"
	"net/http"

	"github.com/evrone/go-clean-template/internal/controller/http/v1/request"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/gofiber/fiber/v2"
)

// @Summary     Show comments
// @Description Show all comments
// @ID          comments
// @Tags  	    comment
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.CommentsList
// @Failure     500 {object} response.Error
// @Router      /comment/comments [get]
func (router *V1) comments(ctx *fiber.Ctx) error {
	commentsHistory, err := router.comment.History(ctx.UserContext())
	if err != nil {
		router.logger.Error(err, "http - v2 - comments")

		return errorResponse(ctx, http.StatusInternalServerError, "database problems")
	}

	return ctx.Status(http.StatusOK).JSON(commentsHistory)
}

// @Summary     Comment
// @Description Comment a entity
// @ID          do-comment
// @Tags  	    comment
// @Accept      json
// @Produce     json
// @Param       request body request.Comment true "Set up comment"
// @Success     200 {object} entity.Comment
// @Failure     400 {object} response.Error
// @Failure     500 {object} response.Error
// @Router      /comments/do-comment [post]
func (router *V1) doComment(ctx *fiber.Ctx) error {
	var body request.Comment

	if err := ctx.BodyParser(&body); err != nil {
		router.logger.Error(err, "http - v1 - doComment")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if err := router.validator.Struct(body); err != nil {
		router.logger.Error(err, "http - v1 - doComment")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	comment, err := router.comment.Comment(
		ctx.UserContext(),
		entity.Comment{
			Text:      body.Text,
			CreatedBy: body.CreatedBy,
		},
		entity.Entity{
			EntityID:   body.EntityID,
			EntityType: body.EntityType,
		},
	)
	if err != nil {
		router.logger.Error(err, "http - v1 - doComment")

		return errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("comment service problems %s", err))
	}

	return ctx.Status(http.StatusOK).JSON(comment)
}
