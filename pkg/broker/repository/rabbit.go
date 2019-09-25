package repository

import (
	"github.com/IamStubborN/calendar/pkg/broker"
	"github.com/streadway/amqp"
)

type brokerRepository struct {
	ch *amqp.Channel
}

func NewBrokerRabbitMQ(ch *amqp.Channel) broker.Repository {
	return &brokerRepository{
		ch: ch,
	}
}

func (br *brokerRepository) Publish(queueName string, data []byte) error {
	q, err := br.ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	return br.ch.Publish(
		q.Name,
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
}

func (br *brokerRepository) Receive(queueName string) ([]byte, error) {
	q, err := br.ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	message, err := br.ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	data := <-message

	return data.Body, err
}
