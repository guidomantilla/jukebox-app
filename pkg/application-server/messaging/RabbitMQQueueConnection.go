package messaging

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

const (
	makeConnectionDelay = 2 * time.Second
)

type DialRabbitMQFunc func(url string) (*amqp.Connection, error)

var _ RabbitMQQueueConnection = (*DefaultRabbitMQQueueConnection)(nil)

type RabbitMQQueueConnection interface {
	Start()
	Close()
	Connect() (*amqp.Connection, *amqp.Channel, *amqp.Queue, error)
}

type DefaultRabbitMQQueueConnection struct {
	rabbitMQUrl              string
	connection               *amqp.Connection
	notifyOnClosedConnection chan *amqp.Error
	channel                  *amqp.Channel
	notifyOnClosedChannel    chan *amqp.Error
	queueName                string
	queue                    amqp.Queue

	isReady                bool
	notifyOnClosingWatcher chan bool
	dialFunc               DialRabbitMQFunc
}

func NewDefaultRabbitMQQueueConnection(queueName string, rabbitMQUrl string, dialFunc DialRabbitMQFunc) *DefaultRabbitMQQueueConnection {

	client := &DefaultRabbitMQQueueConnection{
		rabbitMQUrl:            rabbitMQUrl,
		queueName:              queueName,
		isReady:                false,
		notifyOnClosingWatcher: make(chan bool),
		dialFunc:               dialFunc,
	}

	return client
}

//

func (client *DefaultRabbitMQQueueConnection) Start() {
	go client.watchConnection()
	<-time.After(time.Second)
}

func (client *DefaultRabbitMQQueueConnection) Close() {
	client.notifyOnClosingWatcher <- true
}

func (client *DefaultRabbitMQQueueConnection) Connect() (*amqp.Connection, *amqp.Channel, *amqp.Queue, error) {

	if client.isReady {
		return client.connection, client.channel, &client.queue, nil
	}

	<-time.After(makeConnectionDelay)
	if !client.isReady {
		return nil, nil, nil, fmt.Errorf("unable to connect to %s", client.rabbitMQUrl)
	}

	return client.connection, client.channel, &client.queue, nil
}

//

func (client *DefaultRabbitMQQueueConnection) makeConnection() error {

	var err error
	if client.connection, err = client.dialFunc(client.rabbitMQUrl); err != nil {
		return err
	}

	if client.channel, err = client.connection.Channel(); err != nil {
		return err
	}

	if client.queue, err = client.channel.QueueDeclare(client.queueName, false, false, false, false, nil); err != nil {
		return err
	}

	client.notifyOnClosedConnection = make(chan *amqp.Error, 1)
	client.connection.NotifyClose(client.notifyOnClosedConnection)

	client.notifyOnClosedChannel = make(chan *amqp.Error, 1)
	client.channel.NotifyClose(client.notifyOnClosedChannel)

	client.isReady = true
	return nil
}

func (client *DefaultRabbitMQQueueConnection) watchConnection() {

	zap.L().Info("rabbitmq - connection demon started")
	defer zap.L().Info("rabbitmq - connection demon stopped")

	for {

		if !client.isReady {
			var err error
			if err = client.makeConnection(); err != nil {
				zap.L().Debug("rabbitmq - connection demon - failed to connect. retrying...")
				continue
			}
			zap.L().Info("rabbitmq - connection ready")
		}

		select {

		case <-client.notifyOnClosedConnection:
			client.isReady = false
			zap.L().Info("rabbitmq - connection closed. reconnecting...")
			continue

		case <-client.notifyOnClosedChannel:
			client.isReady = false
			zap.L().Info("rabbitmq - channel closed. recreating...")
			continue

		case <-time.After(5 * time.Second):
			if client.isReady {
				var err error
				if _, err = client.channel.QueueInspect(client.queue.Name); err != nil {
					client.isReady = false
					zap.L().Info("rabbitmq - failed to ping the queue. notifying...")
					continue
				}
			}
			continue

		case <-client.notifyOnClosingWatcher:
			return
		}
	}
}
