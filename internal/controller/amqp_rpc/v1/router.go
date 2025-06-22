package v1

import (
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/evrone/go-clean-template/pkg/rabbitmq/rmq_rpc/server"
	"github.com/go-playground/validator/v10"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(routes map[string]server.CallHandler, c usecase.Comment, l logger.Interface) {
	r := &V1{comment: c, logger: l, validator: validator.New(validator.WithRequiredStructEnabled())}

	{
		routes["v1.getComments"] = r.getComments()
	}
}
