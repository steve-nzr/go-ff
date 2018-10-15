package main

import (
	"flyff/core/net"
	gametimer "flyff/world/service/gameTimer"
	"flyff/world/service/messaging"
	"flyff/world/service/messaging/channel"
	"flyff/world/service/playerstate"

	"flyff/core"
)

func main() {
	core.InitiateDbConnection()
	defer core.CloseDbConnection()

	db := core.GetDbConnection()
	db.AutoMigrate(&core.Character{})

	playerstate.Initialize()
	playerstate.AutoMigrate()

	/////////////
	conn := channel.Initialize()
	defer conn.Close()
	defer channel.Channel.Close()
	go messaging.HandleInternalPackets()
	go messaging.HandleExternalPackets()
	////////

	go gametimer.Timer.Start()

	server := net.Create("0.0.0.0:5400")
	server.OnConnected(messaging.HandleNewClient)
	server.OnDisconnected(messaging.HandleRemoveClient)
	server.OnMessage(messaging.HandleNewMessage)
	server.Start()
}
