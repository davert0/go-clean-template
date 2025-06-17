package v2

import (
	"fmt"
	"net/http"

	"github.com/evrone/go-clean-template/internal/controller/http/v2/request"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/gofiber/fiber/v2"
)

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.TranslationHistory
// @Failure     500 {object} response.Error
// @Router      /translation/history [get]
func (r *V2) comments(ctx *fiber.Ctx) error {
	translationHistory, err := r.t.History(ctx.UserContext())
	if err != nil {
		r.l.Error(err, "http - v2 - comments")

		return errorResponse(ctx, http.StatusInternalServerError, "database problems")
	}

	return ctx.Status(http.StatusOK).JSON(translationHistory)
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body request.Translate true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response.Error
// @Failure     500 {object} response.Error
// @Router      /translation/do-translate [post]
func (r *V2) doComment(ctx *fiber.Ctx) error {
	var body request.Comment

	if err := ctx.BodyParser(&body); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if err := r.v.Struct(body); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	comment, err := r.t.Comment(
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
		r.l.Error(err, "http - v1 - doComment")

		return errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("comment service problems %s", err))
	}

	return ctx.Status(http.StatusOK).JSON(comment)
}
