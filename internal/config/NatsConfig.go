package config

import (
	"github.com/qmdx00/lifecycle"

	appserver "jukebox-app/pkg/application-server"
	"jukebox-app/pkg/environment"
)

const ()

func InitNatsDispatcher(environment environment.Environment) lifecycle.Server {

	//

	natsDispatcher := appserver.BuildNatsMessageDispatcher(appserver.NewDefaultNatsMessageListener())
	return natsDispatcher
}
