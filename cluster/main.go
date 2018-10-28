package main

import (
	"flyff/common/def/resources"
	"flyff/common/feature/inventory/def"
	"flyff/common/service/database"
	"flyff/common/service/dotenv"
	"flyff/common/service/external"
	"flyff/connectionserver/service/connectionmanager"

	"flyff/cluster/packets"
)

func onConnectedHandler(ch <-chan *external.Client) {
	for {
		c := <-ch
		if c == nil {
			continue
		}

		connectionmanager.Add(c)
		c.SendGreetings()
	}
}

func onDisconnectedHandler(ch <-chan *external.Client) {
	for {
		c := <-ch
		if c == nil {
			continue
		}

		connectionmanager.Remove(c)
	}
}

func onMessageHandler(ch <-chan *external.PacketHandler) {
	for {
		p := <-ch
		if p == nil {
			continue
		}

		c := connectionmanager.Get(p.ClientID)
		if c == nil {
			continue
		}

		// Always FFFFFFF
		p.Packet.ReadUInt32()

		protocol := p.Packet.ReadUInt32()

		if protocol == 0xf6 {
			sendPlayerList(c, 0)
			sendWorldAddr(c)
		} else if protocol == 0xf4 {
			var createPlayerPacket packets.CreatePlayer
			createPlayerPacket.Construct(p.Packet)

			var player database.Player
			player.Slot = createPlayerPacket.Slot
			player.Name = createPlayerPacket.Name
			player.Gender = createPlayerPacket.Gender
			player.Position.MapID = 1
			player.Position.Vec.X = 6968.0
			player.Position.Vec.Y = 100.0
			player.Position.Vec.Z = 3328.0
			player.SkinSetID = uint32(createPlayerPacket.SkinSet)
			player.HairID = uint32(createPlayerPacket.HairMeshID)
			player.HairColor = createPlayerPacket.HairColor
			player.FaceID = uint32(createPlayerPacket.FaceID)
			player.JobID = createPlayerPacket.Job
			player.Level = 1
			player.Statistics.Strength = 15
			player.Statistics.Stamina = 15
			player.Statistics.Dexterity = 15
			player.Statistics.Intelligence = 15

			// Start items
			player.Items = append(player.Items, database.Item{
				ItemBase: def.ItemBase{
					Count:    1,
					ItemID:   resources.ItemsByName["II_WEA_SWO_WOODEN"],
					Position: 52,
				},
			})
			player.Items = append(player.Items, database.Item{
				ItemBase: def.ItemBase{
					Count:    1,
					ItemID:   resources.ItemsByName["II_ARM_M_VAG_SUIT01"],
					Position: 44,
				},
			})
			player.Items = append(player.Items, database.Item{
				ItemBase: def.ItemBase{
					Count:    1,
					ItemID:   resources.ItemsByName["II_ARM_M_VAG_GAUNTLET01"],
					Position: 46,
				},
			})
			player.Items = append(player.Items, database.Item{
				ItemBase: def.ItemBase{
					Count:    1,
					ItemID:   resources.ItemsByName["II_ARM_M_VAG_BOOTS01"],
					Position: 47,
				},
			})

			database.Connection.Save(&player)
			sendPlayerList(c, 0)
		} else if protocol == 0xf5 {
			var deletePlayerPacket packets.DeletePlayer
			deletePlayerPacket.Construct(p.Packet)

			database.Connection.Delete(&database.Player{}, deletePlayerPacket.PlayerID)
			sendPlayerList(c, 0)
		} else if protocol == 0xff05 {
			var preJoinPacket packets.PreJoin
			preJoinPacket.Construct(p.Packet)

			c.Send(external.MakePacket(external.PREJOIN))
		}
	}
}

func main() {
	dotenv.Initialize()
	database.Initialize()

	// External ----
	onConnected := make(chan *external.Client)
	onDisconnected := make(chan *external.Client)
	onMessage := make(chan *external.PacketHandler)

	go onConnectedHandler(onConnected)
	go onDisconnectedHandler(onDisconnected)
	go onMessageHandler(onMessage)

	server := external.Create("0.0.0.0:28000")
	server.OnConnected(onConnected)
	server.OnDisconnected(onDisconnected)
	server.OnMessage(onMessage)
	server.Start()
}
