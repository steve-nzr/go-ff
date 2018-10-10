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
	Players map[uint32]*structure.WorldClient
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
	m.Maps[1].Players = make(map[uint32]*structure.WorldClient)

	return m
}

func (m *mapManager) Register(wc *structure.WorldClient) {
	gameMap, ok := m.Maps[wc.Character.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", wc.Character.MapID)
		return
	}

	fmt.Println("New player on map", gameMap.ID)

	addObjPacket := out.MakeAddObj(wc)
	m.SendFrom(wc, &addObjPacket)

	for _, player := range gameMap.Players {
		addObjPacket = out.MakeAddObj(player)
		wc.Send(addObjPacket)
	}

	gameMap.Players[uint32(wc.Character.ID)] = wc
}

func (m *mapManager) Unregister(wc *structure.WorldClient) {
	gameMap, ok := m.Maps[wc.Character.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", wc.Character.MapID)
		return
	}

	fmt.Println("Removing player from map", gameMap.ID)
	delObjPacket := out.MakeDeleteObj(wc)
	m.SendFrom(wc, &delObjPacket)

	delete(gameMap.Players, wc.Character.MapID)
}

func (m *mapManager) SendFrom(wc *structure.WorldClient, p *core.Packet) {
	gameMap, ok := m.Maps[wc.Character.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", wc.Character.MapID)
		return
	}

	for _, player := range gameMap.Players {
		if player.ID == wc.ID {
			continue
		}

		fmt.Println("Sending to", player.Character.Name)
		player.Send(*p)
	}
}
