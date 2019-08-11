package mq

import (
  "github.com/streadway/amqp"
)

// IMessagingClient Defines our interface for connecting, producing and consuming messages.
type IMessagingClient interface {
	ConnectToBroker(connectionString string) error
	Publish(msg []byte, exchangeName, exchangeType, bindingKey string) error
	Subscribe(exchangeName, exchangeType, queueName, bindingKey, consumerName string, handleFunc func(amqp.Delivery)) error
  ExChangeDeclare(name, exchType string) error
  QueueDeclare(name string) (*amqp.Queue, error)
	Close()
}

// MessagingClient Real implementation, encapsulates a pointer to an amqp.Connection
type MessagingClient struct {
	conn *amqp.Connection
}
