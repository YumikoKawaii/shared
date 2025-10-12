package pubsub

import "context"

type Publisher interface {
	Publish(ctx context.Context, topic string, bytes []byte) error
}
