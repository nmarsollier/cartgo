package r_consume

import (
	"time"

	"github.com/golang/glog"
)

func Init() {
	go func() {
		for {
			err := consumeLogout()
			if err != nil {
				glog.Error(err)
			}
			glog.Info("RabbitMQ listenLogout conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := consumeOrders()
			if err != nil {
				glog.Error(err)
			}
			glog.Info("RabbitMQ consumeCart conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			err := consumeOrderPlaced()
			if err != nil {
				glog.Error(err)
			}
			glog.Info("RabbitMQ consumeOrderPlaced conectando en 5 segundos.")
			time.Sleep(5 * time.Second)
		}
	}()

}
