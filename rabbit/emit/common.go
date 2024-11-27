package emit

import (
	"errors"

	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/log"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.New("channel not initialized")

func getChannel(deps ...interface{}) (RabbitChannel, error) {
	for _, o := range deps {
		if ti, ok := o.(RabbitChannel); ok {
			return ti, nil
		}
	}

	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return rabbitChannel{ch: channel}, nil
}
