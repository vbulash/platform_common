package kafka

import (
	"context"

	"github.com/vbulash/platform_common/pkg/client/kafka/consumer"
)

// Consumer ...
type Consumer interface {
	Consume(ctx context.Context, topicName string, handler consumer.Handler) (err error)
	Close() error
}
