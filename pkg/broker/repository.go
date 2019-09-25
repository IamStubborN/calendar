package broker

type Repository interface {
	Publish(queueName string, data []byte) error
	Receive(queueName string) ([]byte, error)
}
