package repository

import (
  "github.com/streadway/amqp"  
)

func (repo *repo) MQPublish(
  msg []byte,
  exchangeName,
  exchangeType,
  bindingKey,
  queueName string) error {
  return repo.MQ.Publish(msg, exchangeName, exchangeType, bindingKey, queueName)
}

func (repo *repo) MQSubscribe(
  exchangeName,
  exchangeType,
  queueName,
  bindingKey,
  consumerName string,
  handleFunc func(amqp.Delivery)) error {
  return repo.MQ.Subscribe(exchangeName, exchangeType, queueName, bindingKey, consumerName, handleFunc)
}
