package usecase

import (
	"context"

	"github.com/IamStubborN/calendar/notifier/pkg/broker"
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
	return br.ch.Publish("", queueName, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
}

func (br *brokerRepository) Receive(ctx context.Context, queueName string) (<-chan string, error) {
	dataCh := make(chan string)
	message, err := br.ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case data := <-message:
				dataCh <- string(data.Body)
			}
		}
	}()

	return dataCh, nil
}
