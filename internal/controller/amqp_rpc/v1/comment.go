package v1

import (
	"github.com/evrone/go-clean-template/pkg/rabbitmq/rmq_rpc/server"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *V1) getComments() server.CallHandler {
	return func(_ *amqp.Delivery) (interface{}, error) {
		//translationHistory, err := r.getUC.GetComments(context.Background())
		// if err != nil {
		// 	r.logger.Error(err, "amqp_rpc - V1 - GetComments")

		// 	return nil, fmt.Errorf("amqp_rpc - V1 - GetComments: %w", err)
		// }

		//return commentsHistory, nil

		return nil, nil
	}
}
