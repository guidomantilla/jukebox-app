package config

import (
	appserver "jukebox-app/pkg/application-server"
	"jukebox-app/pkg/application-server/messaging"
	"jukebox-app/pkg/environment"

	"github.com/qmdx00/lifecycle"
	"go.uber.org/zap"
)

const (
	QUEUE_NAME        = "QUEUE_NAME"
	RABBITMQ_ADDRESS  = "RABBITMQ_ADDRESS"
	RABBITMQ_USERNAME = "RABBITMQ_USERNAME"
	RABBITMQ_PASSWORD = "RABBITMQ_PASSWORD"
)

func InitRabbitMQDispatcher(environment environment.Environment) lifecycle.Server {

	queueName := environment.GetValue(QUEUE_NAME).AsString()
	if queueName == "" {
		zap.L().Fatal("server starting up - error setting up rabbitmq dispatcher: invalid queue name")
	}

	address := environment.GetValue(RABBITMQ_ADDRESS).AsString()
	if address == "" {
		zap.L().Fatal("server starting up - error setting up rabbitmq dispatcher: invalid address")
	}

	username := environment.GetValue(RABBITMQ_USERNAME).AsString()
	password := environment.GetValue(RABBITMQ_PASSWORD).AsString()

	client := messaging.NewDefaultRabbitMQQueueConnection(queueName, username, password, address)
	listener := appserver.NewDefaultRabbitMQMessageListener()

	//

	rabbitMQDispatcher := appserver.BuildRabbitMQMessageDispatcher(client, listener)
	return rabbitMQDispatcher
}
