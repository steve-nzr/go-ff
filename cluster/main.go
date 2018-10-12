package main

import (
	"fmt"

	"flyff/cluster/packets"
	"flyff/core"
	"flyff/core/net"
)

func main() {
	core.InitiateDbConnection()
	defer core.CloseDbConnection()

	db := core.GetDbConnection()
	db.AutoMigrate(&core.Character{})

	server := net.Create("0.0.0.0:28000")
	server.OnConnected(onConnectionInitiated)
	server.OnDisconnected(onConnectionInitiated)
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
	// Always FFFFFFF
	packet.ReadUInt32()

	protocol := packet.ReadUInt32()
	fmt.Printf("New packet with id : 0x%02x\n", protocol)

	if protocol == 0xf6 {
		sendPlayerList(nc, 0)
		sendWorldAddr(nc)
	} else if protocol == 0xf4 {
		var p packets.CreatePlayer
		p.Construct(packet)

		var c core.Character
		c.Slot = p.Slot
		c.Name = p.Name
		c.Gender = p.Gender
		c.MapID = 1
		c.PosX = 6968.0
		c.PosY = 100.0
		c.PosZ = 3328.0
		c.SkinSetID = uint32(p.SkinSet)
		c.HairID = uint32(p.HairMeshID)
		c.HairColor = p.HairColor
		c.FaceID = uint32(p.FaceID)
		c.JobID = p.Job
		c.Level = 1
		c.Strength = 15
		c.Stamina = 15
		c.Dexterity = 15
		c.Intelligence = 15

		db := core.GetDbConnection()
		db.Save(&c)

		sendPlayerList(nc, 0)
	} else if protocol == 0xf5 {
		var p packets.DeletePlayer
		p.Construct(packet)

		db := core.GetDbConnection()
		db.Delete(&core.Character{}, p.CharacterID)

		sendPlayerList(nc, 0)
	} else if protocol == 0xff05 {
		var p packets.PreJoin
		p.Construct(packet)

		nc.Send(net.MakePacket(net.PREJOIN))
	}
}
