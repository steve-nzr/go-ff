package main

import (
	"fmt"

	"flyff/core"
)

func main() {
	core.StartNetServer(core.NetServerConfig{
		Host:                "0.0.0.0",
		Port:                "23000",
		Type:                core.NetServerTCP,
		ConnectionInitiated: onConnectionInitiated,
		ConnectionMessage:   onConnectionMessage})
}

func onConnectionInitiated(c *core.NetClient) {
	loginClient{c}.SendGreetings()
}

func onConnectionMessage(c *core.NetClient, packet *core.Packet) {
	lc := loginClient{c}

	protocol := packet.ReadUInt32()
	fmt.Printf("New packet with id : 0x%02x\n", protocol)

	if protocol == 0xfc {
		fmt.Println("Login request")

		lc.sendServerList()
	}
}
