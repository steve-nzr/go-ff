package cache

import (
	"log"
)

// FindByNetID return the player who has this clientID
func FindByNetID(id uint32) *Player {
	player := new(Player)
	err := Connection.Where("net_client_id = ?", id).First(player).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return player
}

// FindIDAround returns each player'ID around the given one
func FindIDAround(p *Player) []uint32 {
	var IDlist []uint32
	err := Connection.Model(&Player{}).Where("posit_map_id = ?", p.Position.MapID).Pluck("net_client_id", &IDlist).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return IDlist
}

// FindAround returns each player structure around the given one
func FindAround(p *Player) []Player {
	var playerList []Player
	err := Connection.Where("posit_map_id = ?", p.Position.MapID).Find(&playerList).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return playerList
}
