package app

import (
	"fmt"

	"github.com/IamStubborN/calendar/reminder/config"
	"github.com/IamStubborN/calendar/reminder/pkg/broker"
	"github.com/IamStubborN/calendar/reminder/pkg/broker/usecase"
	"github.com/IamStubborN/calendar/reminder/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func initializeBroker(cfg *config.Config, logger logger.UseCase) broker.Repository {
	conn, err := amqp.Dial(cfg.Broker.DSN)
	if err != nil {
		fmt.Print(1)
		logger.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Fatal(err)
	}

	_, err = ch.QueueDeclare(
		cfg.Broker.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Fatal(err)
	}

	br := usecase.NewBrokerRabbitMQ(ch)
	if err != nil {
		logrus.Fatalln(err)
	}

	return br
}
