package emit

import (
	"github.com/golang/mock/gomock"
)

func DefaultRabbitChannel(ctrl *gomock.Controller, times int) *MockRabbitChannel {
	channel := NewMockRabbitChannel(ctrl)
	channel.EXPECT().ExchangeDeclare(gomock.Any(), gomock.Any()).Return(nil).Times(times)
	return channel
}
