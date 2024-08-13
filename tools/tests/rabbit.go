package tests

import (
	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/rabbit/r_emit"
)

func MockRabbitChannel(ctrl *gomock.Controller, times int) *r_emit.MockRabbitChannel {
	channel := r_emit.NewMockRabbitChannel(ctrl)
	channel.EXPECT().ExchangeDeclare(gomock.Any(), gomock.Any()).Return(nil).Times(times)
	return channel
}
