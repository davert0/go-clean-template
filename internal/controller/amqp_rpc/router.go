package v1

import (
	v1 "github.com/evrone/go-clean-template/internal/controller/amqp_rpc/v1"
	"github.com/evrone/go-clean-template/internal/usecase/comment/create"
	"github.com/evrone/go-clean-template/internal/usecase/comment/get"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/evrone/go-clean-template/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(c create.UseCase, g get.UseCase, l logger.Interface) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)

	{
		v1.NewCommentRoutes(routes, c, g, l)
	}

	return routes
}
