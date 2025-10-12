package pubsub

import "context"

type HandleMessageFn func(bytes []byte) error

type Subscriber interface {
	Consume(ctx context.Context, topic string, fn HandleMessageFn)
	Close(ctx context.Context) error
}
