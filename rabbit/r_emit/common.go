package r_emit

import (
	"errors"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.New("channel not initialized")

func getChannel(ctx ...interface{}) (RabbitChannel, error) {
	for _, o := range ctx {
		if ti, ok := o.(RabbitChannel); ok {
			return ti, nil
		}
	}

	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return rabbitChannel{ch: channel}, nil
}
