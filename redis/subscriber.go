package redis

import (
	"context"

	"github.com/YumikoKawaii/shared/pubsub"
	v9 "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type subscriberImpl struct {
	redisClient *v9.Client
	subscriber  *v9.PubSub
}

func NewSubscriber(client *v9.Client) pubsub.Subscriber {
	return &subscriberImpl{
		redisClient: client,
	}
}

func (c *subscriberImpl) Consume(ctx context.Context, topic string, fn pubsub.HandleMessageFn) {

	subscriber := c.redisClient.Subscribe(ctx, topic)
	c.subscriber = subscriber
	for {
		message, err := c.subscriber.ReceiveMessage(ctx)
		if err != nil {
			logrus.Errorf("error while receiving message: %s", err.Error())
			continue
		}

		bytes := []byte(message.Payload)

		if err := fn(bytes); err != nil {
			logrus.Errorf("error handling message: %s", err.Error())
		}
	}
}

func (c *subscriberImpl) Close(ctx context.Context) error {
	return c.subscriber.Close()
}
