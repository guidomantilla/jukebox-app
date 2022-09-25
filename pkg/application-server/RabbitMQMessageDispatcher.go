package applicationserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/qmdx00/lifecycle"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"

	"jukebox-app/pkg/application-server/messaging"
)

//

type RabbitMQMessageListener interface {
	OnMessage(message *amqp.Delivery) error
}

type DefaultRabbitMQMessageListener struct {
}

func NewDefaultRabbitMQMessageListener() *DefaultRabbitMQMessageListener {
	return &DefaultRabbitMQMessageListener{}
}

func (listener *DefaultRabbitMQMessageListener) OnMessage(message *amqp.Delivery) error {
	zap.L().Info(fmt.Sprintf("Received a message: %s", message.Body))
	<-time.After(5 * time.Second)
	return nil
}

//

var _ lifecycle.Server = (*RabbitMQMessageDispatcher)(nil)

type RabbitMQMessageDispatcher struct {
	rabbitmqConnection   messaging.RabbitMQQueueConnection
	rabbitmqListener     RabbitMQMessageListener
	ctx                  context.Context
	receivedMessagesChan <-chan amqp.Delivery
}

func BuildRabbitMQMessageDispatcher(rabbitmqConnection messaging.RabbitMQQueueConnection, rabbitmqListener RabbitMQMessageListener) lifecycle.Server {
	return &RabbitMQMessageDispatcher{
		rabbitmqConnection: rabbitmqConnection,
		rabbitmqListener:   rabbitmqListener,
	}
}

func (server *RabbitMQMessageDispatcher) Run(ctx context.Context) error {

	server.ctx = ctx
	info, _ := lifecycle.FromContext(ctx)
	zap.L().Info(fmt.Sprintf("server starting up - starting rabbitmq dispatcher %s, v.%s", info.Name(), info.Version()))

	server.rabbitmqConnection.Start()

	var err error
	//var connection *amqp.Connection
	var channel *amqp.Channel
	var queue *amqp.Queue

	if _, channel, queue, err = server.rabbitmqConnection.Connect(); err != nil {
		zap.L().Error(fmt.Sprintf("server starting up - rabbitmq dispatcher - error: %s", err.Error()))
		return err
	}

	if server.receivedMessagesChan, err = channel.Consume(queue.Name, "", true, false, false, false, nil); err != nil {
		zap.L().Error(fmt.Sprintf("server starting up - rabbitmq dispatcher - error: %s", err.Error()))
		return err
	}

	if err = server.ListenAndDispatch(); err != nil && err != http.ErrServerClosed {
		zap.L().Error(fmt.Sprintf("server starting up - rabbitmq dispatcher - error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *RabbitMQMessageDispatcher) ListenAndDispatch() error {

	for {
		select {
		case <-server.ctx.Done():
			return nil
		case message := <-server.receivedMessagesChan:
			go server.Dispatch(&message)
		}
	}
}

func (server *RabbitMQMessageDispatcher) Dispatch(message *amqp.Delivery) {

	var err error
	if err = server.rabbitmqListener.OnMessage(message); err != nil {
		zap.L().Error(fmt.Sprintf("rabbitmq listener - error: %s, message: %s", err.Error(), message.Body))
	}
}

func (server *RabbitMQMessageDispatcher) Stop(ctx context.Context) error {

	info, _ := lifecycle.FromContext(ctx)
	zap.L().Info(fmt.Sprintf("server shutting down - stopping rabbitmq dispatcher %s, v.%s", info.Name(), info.Version()))

	server.rabbitmqConnection.Close()

	zap.L().Info("server shutting down - rabbitmq dispatcher stopped")
	return nil
}
