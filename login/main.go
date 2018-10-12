package main

import (
	"fmt"

	"flyff/core/net"
)

func main() {
	server := net.Create("0.0.0.0:23000")
	server.OnConnected(onConnectionInitiated)
	server.OnDisconnected(onConnectionClosed)
	server.OnMessage(onConnectionMessage)
	server.Start()
}

func onConnectionInitiated(nc *net.Client) {
	fmt.Println("Client", nc.ID, "connected")
	nc.SendGreetings()
}

func onConnectionClosed(nc *net.Client) {
	fmt.Println("Client", nc.ID, "disconnected")
}

func onConnectionMessage(nc *net.Client, packet *net.Packet) {
	switch packet.ReadUInt32() {
	case 0xfc:
		{
			sendServerList(nc)
		}
	}
}
