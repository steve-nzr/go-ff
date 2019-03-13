package cache

import (
	"log"
)

// FindByNetID return the player who has this clientID
func FindByNetID(id uint32) *Player {
	player := new(Player)
	err := Connection.Preload("Inventory").Where("net_client_id = ?", id).First(player).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return player
}

// FindIDAroundOnly returns each player'ID around the given one
func FindIDAroundOnly(p *Player) []uint32 {
	var IDlist []uint32
	err := Connection.Model(&Player{}).Where("posit_map_id = ? AND net_client_id != ?", p.Position.MapID, p.NetClientID).Pluck("net_client_id", &IDlist).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return IDlist
}

// FindIDAround returns each player'ID around the given one including itself
func FindIDAround(p *Player) []uint32 {
	var IDlist []uint32
	err := Connection.Model(&Player{}).Where("posit_map_id = ?", p.Position.MapID).Pluck("net_client_id", &IDlist).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return IDlist
}

// FindAround returns each player structure around the given one including itself
func FindAround(p *Player) []Player {
	var playerList []Player
	err := Connection.Preload("Inventory").Where("posit_map_id = ?", p.Position.MapID).Find(&playerList).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return playerList
}

// FindAroundOnly returns each player structure around the given one excluding itself
func FindAroundOnly(p *Player) []Player {
	var playerList []Player
	err := Connection.Preload("Inventory").Where("posit_map_id = ? AND net_client_id != ?", p.Position.MapID, p.NetClientID).Find(&playerList).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return playerList
}
