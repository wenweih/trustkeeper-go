package repository

import (
	"github.com/streadway/amqp"
)

const (
	// BindKeyBitcoincore BindKeyBitcoincore
	BindKeyBitcoincore string = "bitcoincore"
	// BindKeyEthereum BindKeyEthereum
	BindKeyEthereum string = "ethereum"
	// ExchangeNameForBitcoincoreBestBlock ExchangeNameForBitcoincoreBestBlock
	ExchangeNameForBitcoincoreBestBlock string = "bitcoincorebestblock"
	// ExchangeNameForEthereumBestBlock ExchangeNameForEthereumBestBlock
	ExchangeNameForEthereumBestBlock string = "ethereumbestblock"

	// QueueNameForBitcoincoreTraceTx QueueNameForBitcoincoreQueue
	QueueNameForBitcoincoreTraceTx string = "queue_bitcoincore_trace_tx"
	// QueueNameForEthereumTraceTx QueueNameForEthereumQueue
	QueueNameForEthereumTraceTx string = "queue_ethereum_trace_tx"
	// QueueNameForBitcoincoreUpdateTx QueueNameForBitcoincoreUpdateTx
	QueueNameForBitcoincoreUpdateTx string = "queue_bitcoincore_update_tx"
	// QueueNameForEthereumUpdateTx QueueNameForEthereumUpdateTx
	QueueNameForEthereumUpdateTx string = "queue_ethereum_update_tx"
)

func (repo *repo) DeclareExChange(exchangeName, exchType string) error {
	return repo.MQ.ExChangeDeclare(exchangeName, exchType)
}

func (repo *repo) DeclareQueue(queueName string) (*amqp.Queue, error) {
	return repo.MQ.QueueDeclare(queueName)
}

func (repo *repo) MQPublish(
	msg []byte,
	exchangeName,
	exchangeType,
	bindingKey string) error {
	return repo.MQ.Publish(msg, exchangeName, exchangeType, bindingKey)
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
