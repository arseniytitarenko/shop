package out

import (
	"context"
)

type BrokerPublisher interface {
	Publish(ctx context.Context, messageType string, payload string) error
}
