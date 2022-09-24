package application_server

import (
	"context"
	"fmt"

	"github.com/qmdx00/lifecycle"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"

	"jukebox-app/pkg/application-server/messaging"
)

var _ lifecycle.Server = (*RabbitMQMessageDispatcher)(nil)

type RabbitMQMessageDispatcher struct {
	rabbitmqConnection *messaging.DefaultRabbitMQQueueConnection
	forever            chan int
}

func BuildRabbitMQMessageDispatcher(rabbitmqConnection *messaging.DefaultRabbitMQQueueConnection) lifecycle.Server {
	return &RabbitMQMessageDispatcher{
		rabbitmqConnection: rabbitmqConnection,
		forever:            make(chan int),
	}
}

func (server *RabbitMQMessageDispatcher) Run(ctx context.Context) error {

	info, _ := lifecycle.FromContext(ctx)
	zap.L().Info(fmt.Sprintf("server starting up - starting rabbitmq listener %s, v.%s", info.Name(), info.Version()))

	server.rabbitmqConnection.Start()

	var err error
	//var connection *amqp.Connection
	var channel *amqp.Channel
	var queue *amqp.Queue

	if _, channel, queue, err = server.rabbitmqConnection.Connect(); err != nil {
		zap.L().Error(fmt.Sprintf("server starting up - error: %s", err.Error()))
		return err
	}

	var receivedMessagesChan <-chan amqp.Delivery
	if receivedMessagesChan, err = channel.Consume(queue.Name, "", true, false, false, false, nil); err != nil {
		zap.L().Error(fmt.Sprintf("server starting up - error: %s", err.Error()))
		return err
	}

	go func() {
		for d := range receivedMessagesChan {
			zap.L().Info(fmt.Sprintf("Received a message: %s", d.Body))
		}
	}()
	<-server.forever

	return nil
}

func (server *RabbitMQMessageDispatcher) Stop(ctx context.Context) error {

	info, _ := lifecycle.FromContext(ctx)
	zap.L().Info(fmt.Sprintf("server shutting down - stopping rabbitmq listener %s, v.%s", info.Name(), info.Version()))

	close(server.forever)
	server.rabbitmqConnection.Close()

	zap.L().Info("server shutting down - rabbitmq listener stopped")
	return nil
}
