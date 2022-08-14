package main

import (
	"fmt"
	"jukebox-app/examples/structural-patterns/adapter/machines"
)

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com machines.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
