package tests

import (
	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/rabbit/emit"
)

func MockRabbitChannel(ctrl *gomock.Controller, times int) *emit.MockRabbitChannel {
	channel := emit.NewMockRabbitChannel(ctrl)
	channel.EXPECT().ExchangeDeclare(gomock.Any(), gomock.Any()).Return(nil).Times(times)
	return channel
}
