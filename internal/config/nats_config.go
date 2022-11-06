package config

import (
	"github.com/qmdx00/lifecycle"

	appserver "jukebox-app/pkg/app"
	"jukebox-app/pkg/environment"
)

const ()

func InitNatsDispatcher(environment environment.Environment) lifecycle.Server {

	//

	natsDispatcher := appserver.BuildNatsMessageDispatcher(appserver.NewDefaultNatsMessageListener())
	return natsDispatcher
}
