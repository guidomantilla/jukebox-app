package config

import (
	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	"github.com/qmdx00/lifecycle"

	appserver "jukebox-app/pkg/app"
)

const ()

func InitNatsDispatcher(environment environment.Environment) lifecycle.Server {

	//

	natsDispatcher := appserver.BuildNatsMessageDispatcher(appserver.NewDefaultNatsMessageListener())
	return natsDispatcher
}
