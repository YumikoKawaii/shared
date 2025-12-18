package redis

import (
	"context"

	"github.com/YumikoKawaii/shared/pubsub"
	v9 "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type publisherImpl struct {
	redisClient *v9.Client
}

func NewPublisher(client *v9.Client) pubsub.Publisher {
	return &publisherImpl{
		redisClient: client,
	}
}

func (p *publisherImpl) Publish(ctx context.Context, topic string, bytes []byte) error {

	if err := p.redisClient.Publish(ctx, topic, bytes).Err(); err != nil {
		logrus.Errorf("error while publish message to topic %s: %s", topic, err.Error())
	}
	logrus.Infof("Event was sent to topic %s", topic)
	return nil

}
