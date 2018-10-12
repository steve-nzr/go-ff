package gamemap

import (
	"flyff/core/net"
	"flyff/world/game/structure"
	"flyff/world/packets/out"
	"fmt"
	"log"
)

type manager struct {
	Maps map[uint32]*gameMap
}

var Manager = initialize()

func initialize() *manager {
	m := new(manager)
	m.Maps = make(map[uint32]*gameMap)

	m.Maps[1] = new(gameMap)
	m.Maps[1].ID = 1
	m.Maps[1].Players = make(map[uint32]*structure.PlayerEntity)

	return m
}

func (m *manager) Update(time int64) {
	for _, gm := range m.Maps {
		gm.Update(time)
	}
}

func (m *manager) Register(wc *structure.WorldClient) {
	gameMap, ok := m.Maps[wc.PlayerEntity.Position.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", wc.PlayerEntity.Position.MapID)
		return
	}

	fmt.Println("New player on map", gameMap.ID)

	addObjPacket := out.MakeAddObj(wc.PlayerEntity)
	m.SendFrom(wc, &addObjPacket)

	for _, player := range gameMap.Players {
		addObjPacket = out.MakeAddObj(player)
		wc.Send(addObjPacket)
	}

	gameMap.Players[uint32(wc.PlayerEntity.ID)] = wc.PlayerEntity
}

func (m *manager) Unregister(wc *structure.WorldClient) {
	gameMap, ok := m.Maps[wc.PlayerEntity.Position.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", wc.PlayerEntity.Position.MapID)
		return
	}

	fmt.Println("Removing player from map", gameMap.ID)
	delObjPacket := out.MakeDeleteObj(wc.PlayerEntity)
	m.SendFrom(wc, &delObjPacket)

	delete(gameMap.Players, wc.PlayerEntity.Position.MapID)
}

func (m *manager) SendFrom(wc *structure.WorldClient, p *net.Packet) {
	gameMap, ok := m.Maps[wc.PlayerEntity.Position.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", wc.PlayerEntity.Position.MapID)
		return
	}

	for _, player := range gameMap.Players {
		if player.ID == wc.PlayerEntity.ID {
			continue
		}

		fmt.Println("Sending to", player.Name)
		player.WorldClient.Send(*p)
	}
}
