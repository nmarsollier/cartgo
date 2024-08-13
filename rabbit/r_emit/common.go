package r_emit

import (
	"errors"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/streadway/amqp"
)

// ErrChannelNotInitialized Rabbit channel could not be initialized
var ErrChannelNotInitialized = errors.New("Channel not initialized")

func getChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	if ch == nil {
		return nil, ErrChannelNotInitialized
	}
	return ch, nil
}
