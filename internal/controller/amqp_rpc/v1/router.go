package v1

import (
	"github.com/evrone/go-clean-template/internal/usecase/comment/create"
	"github.com/evrone/go-clean-template/internal/usecase/comment/history"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/evrone/go-clean-template/pkg/rabbitmq/rmq_rpc/server"
	"github.com/go-playground/validator/v10"
)

// NewCommentRoutes -.
func NewCommentRoutes(routes map[string]server.CallHandler, c create.UseCase, h history.UseCase, l logger.Interface) {
	r := &V1{createUC: c, getUC: h, logger: l, validator: validator.New(validator.WithRequiredStructEnabled())}

	{
		routes["v1.getHistory"] = r.getComments()
	}
}
