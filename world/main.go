package main

import (
	"flyff/world/game/structure"
	"flyff/world/packets/in"
	"flyff/world/service/gamemap"
	"fmt"
	"sync"

	"flyff/core"
)

type WorldClients map[uint32]*structure.WorldClient

var clients = make(WorldClients)
var clientsMutex sync.RWMutex

func main() {
	core.InitiateDbConnection()
	defer core.CloseDbConnection()

	db := core.GetDbConnection()
	db.AutoMigrate(&core.Character{})

	core.StartNetServer(core.NetServerConfig{
		Host:                "0.0.0.0",
		Port:                "5400",
		Type:                core.NetServerTCP,
		ConnectionClosed:    onConnectionClosed,
		ConnectionInitiated: onConnectionInitiated,
		ConnectionMessage:   onConnectionMessage})
}

func onConnectionClosed(c *core.NetClient) {
	clientsMutex.Lock()
	wc := clients[c.ID]
	defer clientsMutex.Unlock()

	gamemap.Manager.Unregister((*structure.WorldClient)(wc))

	delete(clients, c.ID)
}

func onConnectionInitiated(c *core.NetClient) {
	wc := new(structure.WorldClient)
	wc.NetClient = c

	clientsMutex.Lock()
	clients[c.ID] = wc
	clientsMutex.Unlock()

	wc.SendGreetings()
}

func onConnectionMessage(nc *core.NetClient, packet *core.Packet) {
	clientsMutex.Lock()
	wc := clients[nc.ID]
	clientsMutex.Unlock()
	if wc == nil {
		return
	}

	// Always FFFFFFF
	packet.ReadUInt32()

	protocol := packet.ReadUInt32()
	//fmt.Printf("New packet with id : 0x%02x\n", protocol)

	if protocol == 0xff00 {
		in.Join(wc, packet)
	} else if protocol == 0xffffff00 {
		packet.ReadUInt8()
		snapshotProtocol := packet.ReadUInt16()
		if snapshotProtocol == 0x00C1 {
			fmt.Println("Destpos")
		}
	} else if protocol == 0x00FF0000 {
		in.Chat(wc, packet)
	}
}
