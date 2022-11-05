package applicationserver

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/qmdx00/lifecycle"
	"go.uber.org/zap"
)

//

var _ NatsMessageListener = (*DefaultNatsMessageListener)(nil)

type NatsMessageListener interface {
	OnMessage(message *nats.Msg) error
}

type DefaultNatsMessageListener struct {
}

func NewDefaultNatsMessageListener() *DefaultNatsMessageListener {
	return &DefaultNatsMessageListener{}
}

func (listener *DefaultNatsMessageListener) OnMessage(message *nats.Msg) error {
	zap.L().Info(fmt.Sprintf("Received a message: %s", message.Data))
	<-time.After(5 * time.Second)
	return nil
}

//

var _ lifecycle.Server = (*NatsMessageDispatcher)(nil)

type NatsMessageDispatcher struct {
	natsConnection       *nats.Conn
	natsListener         NatsMessageListener
	ctx                  context.Context
	receivedMessagesChan chan *nats.Msg
}

func BuildNatsMessageDispatcher(natsListener NatsMessageListener) lifecycle.Server {
	return &NatsMessageDispatcher{
		natsListener:         natsListener,
		receivedMessagesChan: make(chan *nats.Msg),
	}
}

func (server *NatsMessageDispatcher) Run(ctx context.Context) error {

	server.ctx = ctx
	info, _ := lifecycle.FromContext(ctx)
	zap.L().Info(fmt.Sprintf("server starting up - starting nats dispatcher %s, v.%s", info.Name(), info.Version()))

	var err error

	if server.natsConnection, err = nats.Connect("nats://192.168.0.101:4222"); err != nil {
		zap.L().Error(fmt.Sprintf("server starting up - nats dispatcher - error: %s", err.Error()))
		return err
	}

	if _, err = server.natsConnection.ChanQueueSubscribe("my-subject", "my-queue", server.receivedMessagesChan); err != nil {
		zap.L().Error(fmt.Sprintf("server starting up - nats dispatcher - error: %s", err.Error()))
		return err
	}
	if err = server.ListenAndDispatch(); err != nil {
		zap.L().Error(fmt.Sprintf("server starting up - nats dispatcher - error: %s", err.Error()))
		return err
	}

	return nil
}

func (server *NatsMessageDispatcher) ListenAndDispatch() error {

	for {
		select {
		case <-server.ctx.Done():
			return nil
		case message := <-server.receivedMessagesChan:
			go server.Dispatch(message)
		}
	}
}

func (server *NatsMessageDispatcher) Dispatch(message *nats.Msg) {

	var err error
	if err = server.natsListener.OnMessage(message); err != nil {
		zap.L().Error(fmt.Sprintf("nats listener - error: %s, message: %s", err.Error(), message.Data))
	}
}

func (server *NatsMessageDispatcher) Stop(ctx context.Context) error {

	info, _ := lifecycle.FromContext(ctx)
	zap.L().Info(fmt.Sprintf("server shutting down - stopping nats dispatcher %s, v.%s", info.Name(), info.Version()))

	server.natsConnection.Close()

	zap.L().Info("server shutting down - nats dispatcher stopped")
	return nil
}
