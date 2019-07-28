package mq

import (
	"errors"
	"strings"
	"github.com/streadway/amqp"
)

// ConnectToBroker connect to rabbitmq broker
func (m *MessagingClient) ConnectToBroker(connectionString string) error {
	if connectionString == "" {
		return errors.New("connectionString empty")
	}

	var err error
	m.conn, err = amqp.Dial(connectionString)
	if err != nil {
		return errors.New("fail to connect rabbitmq broker")
	}
	return nil
}

// Publish publishes a message to the named exchange
func (m *MessagingClient) Publish(body []byte, exchangeName, exchangeType, bindingKey, queueName string) error {
	if m.conn == nil {
		return errors.New("Publish error: conn is nil")
	}

	ch, err := m.conn.Channel() // Get a channel from the connection
	if err != nil {
		return nil
	}
	defer ch.Close()

	if err := ExChangeDeclare(exchangeName, exchangeType, ch); err != nil {
		return err
	}
	if _, err := QueueDeclare(queueName, ch); err != nil {
		return err
	}
	err = ch.Publish(
		exchangeName,
		bindingKey,
		false,
		false,
		amqp.Publishing{
			Body: body,
		})
	return err
}

// Subscribe register a handler function for a given exchange
func (m *MessagingClient) Subscribe(exchangeName, exchangeType, queueName, bindingKey, consumerName string, handleFunc func(amqp.Delivery)) error {
	ch, err := m.conn.Channel()
	if err != nil {
		return errors.New(strings.Join([]string{"Subscribe channel error: "}, err.Error()))
	}
	ExChangeDeclare(exchangeName, exchangeType, ch)
	queue, err := QueueDeclare(queueName, ch)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		queue.Name,
		bindingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return errors.New(strings.Join([]string{"Subscribe queue bind error: ", err.Error()}, ""))
	}
	msgs, err := ch.Consume(
		queue.Name,
		consumerName,
		true,
		false,
		false,
		false,
		nil,
	)
	go consumeLoop(msgs, handleFunc)
	return nil
}

// Close closes the connection to the AMQP-broker, if available
func (m *MessagingClient) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func consumeLoop(deliveries <-chan amqp.Delivery, handlerFunc func(d amqp.Delivery)) {
	for d := range deliveries {
		// Invoke the handlerFunc func we passed as parameter.
		handlerFunc(d)
	}
}

// ExChangeDeclare 声明 rabbitmq exchange
func ExChangeDeclare(name, exchType string, ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		name,     // name
		exchType, // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	return failOnError(err, "Failed to declare an exchange")
}

// QueueDeclare 声明 queue
func QueueDeclare(name string, ch *amqp.Channel) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err := failOnError(err, "Failed to declare a queue"); err != nil {
		return nil, err
	}
	return &q, nil

}

// failOnError RabbitMQ 错误
func failOnError(err error, msg string) error {
	if err != nil {
		return errors.New(strings.Join([]string{msg, err.Error()}, ""))
	}
	return nil
}
