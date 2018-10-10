package mapmanager

import (
	"flyff/core"
	"flyff/world/game/structure"
	"flyff/world/packets/out"
	"fmt"
	"log"
)

type GameMap struct {
	ID      uint32
	Players map[uint32]*structure.PlayerEntity
}

type mapManager struct {
	Maps map[uint32]*GameMap
}

var Manager = initialize()

func initialize() *mapManager {
	m := new(mapManager)
	m.Maps = make(map[uint32]*GameMap)

	m.Maps[1] = new(GameMap)
	m.Maps[1].ID = 1
	m.Maps[1].Players = make(map[uint32]*structure.PlayerEntity)

	return m
}

func (m *mapManager) Register(wc *structure.WorldClient) {
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

func (m *mapManager) Unregister(wc *structure.WorldClient) {
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

func (m *mapManager) SendFrom(wc *structure.WorldClient, p *core.Packet) {
	gameMap, ok := m.Maps[wc.PlayerEntity.Position.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", wc.PlayerEntity.Position.MapID)
		return
	}

	for _, player := range gameMap.Players {
		if player.ID == wc.ID {
			continue
		}

		fmt.Println("Sending to", player.Name)
		player.WorldClient.Send(*p)
	}
}
