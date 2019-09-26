package app

import (
	"github.com/IamStubborN/calendar/config"
	"github.com/IamStubborN/calendar/pkg/broker"
	"github.com/IamStubborN/calendar/pkg/broker/repository"
	"github.com/IamStubborN/calendar/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func initializeBroker(cfg *config.Config, logger logger.Repository) broker.Repository {
	conn, err := amqp.Dial(cfg.Broker.DSN)
	if err != nil {
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

	br := repository.NewBrokerRabbitMQ(ch)
	if err != nil {
		logrus.Fatalln(err)
	}

	return br
}
