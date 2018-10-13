package main

import (
	"flyff/core/net"
	"flyff/world/entities"
	"flyff/world/packets/in"
	gametimer "flyff/world/service/gameTimer"
	"flyff/world/service/gamemap"
	"sync"

	"flyff/core"
)

type WorldClients map[uint32]*entities.PlayerEntity

var clients = make(WorldClients)
var clientsMutex sync.RWMutex

func main() {
	core.InitiateDbConnection()
	defer core.CloseDbConnection()

	db := core.GetDbConnection()
	db.AutoMigrate(&core.Character{})

	go gametimer.Timer.Start()

	server := net.Create("0.0.0.0:5400")
	server.OnConnected(onConnectionInitiated)
	server.OnDisconnected(onConnectionClosed)
	server.OnMessage(onConnectionMessage)
	server.Start()
}

func onConnectionClosed(c *net.Client) {
	clientsMutex.RLock()
	pe := clients[c.ID]
	clientsMutex.RUnlock()

	gamemap.Manager.Unregister(pe)

	clientsMutex.Lock()
	delete(clients, c.ID)
	clientsMutex.Unlock()
}

func onConnectionInitiated(c *net.Client) {
	pe := new(entities.PlayerEntity)
	pe.Client = c

	clientsMutex.RLock()
	clients[c.ID] = pe
	clientsMutex.RUnlock()

	pe.Client.SendGreetings()
}

func onConnectionMessage(nc *net.Client, packet *net.Packet) {
	clientsMutex.RLock()
	pe := clients[nc.ID]
	clientsMutex.RUnlock()
	if pe == nil {
		return
	}

	// Always FFFFFFF
	packet.ReadUInt32()

	protocol := packet.ReadUInt32()

	switch protocol {
	case 0xff00:
		{
			in.Join(pe, packet)
		}
	case 0xffffff00:
		{
			packet.ReadUInt8()
			snapshotProtocol := packet.ReadUInt16()
			if snapshotProtocol == 0x00C1 {
				in.DestPos(pe, packet)
			}
		}
	case 0x00FF0000:
		{
			in.Chat(pe, packet)
		}
	}
}
