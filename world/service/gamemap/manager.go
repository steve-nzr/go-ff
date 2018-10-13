package gamemap

import (
	"flyff/core/net"
	"flyff/world/entities"
	"flyff/world/feature/movement"
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

	gm := new(gameMap)
	gm.ID = 1
	gm.Players = make(map[uint32]*entities.PlayerEntity)
	gm.UpdatableSystems = append(gm.UpdatableSystems, new(movement.System))
	m.Maps[gm.ID] = gm

	return m
}

func (m *manager) Update(time int64) {
	for _, gm := range m.Maps {
		gm.Update(time)
	}
}

func (m *manager) Register(pe *entities.PlayerEntity) {
	gameMap, ok := m.Maps[pe.Position.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", pe.Position.MapID)
		return
	}

	fmt.Println("New player on map", gameMap.ID)

	addObjPacket := out.MakeAddObj(pe)
	m.SendFrom(pe, &addObjPacket)

	for _, player := range gameMap.Players {
		addObjPacket = out.MakeAddObj(player)
		pe.Client.Send(addObjPacket)
	}

	gameMap.Players[uint32(pe.ID)] = pe
}

func (m *manager) Unregister(pe *entities.PlayerEntity) {
	gameMap, ok := m.Maps[pe.Position.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", pe.Position.MapID)
		return
	}

	fmt.Println("Removing player from map", gameMap.ID)
	delObjPacket := out.MakeDeleteObj(pe)
	m.SendFrom(pe, &delObjPacket)

	delete(gameMap.Players, pe.Position.MapID)
}

func (m *manager) SendFrom(pe *entities.PlayerEntity, p *net.Packet) {
	gameMap, ok := m.Maps[pe.Position.MapID]
	if gameMap == nil || ok == false {
		log.Fatalln("GameMap not found", pe.Position.MapID)
		return
	}

	for _, player := range gameMap.Players {
		if player.ID == pe.ID {
			continue
		}

		fmt.Println("Sending to", player.Name)
		player.Client.Send(*p)
	}
}
