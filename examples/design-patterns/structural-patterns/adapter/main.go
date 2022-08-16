package main

import (
	"jukebox-app/examples/design-patterns/structural-patterns/adapter/adapters"
	"jukebox-app/examples/design-patterns/structural-patterns/adapter/machines"
)

func main() {

	client := &Client{}
	mac := &machines.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &machines.Windows{}
	windowsMachineAdapter := &adapters.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
