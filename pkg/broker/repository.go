package broker

import "context"

type Repository interface {
	Publish(queueName string, data []byte) error
	Receive(ctx context.Context, queueName string) (<-chan string, error)
}
