package structure

import (
	"flyff/world/entities"
	"flyff/world/feature"
	"sync"
)

// GameMap represents a Map and it's components
type GameMap struct {
	ID               uint32
	PlayersMutex     sync.RWMutex
	Players          map[uint32]*entities.PlayerEntity
	UpdatableSystems []feature.IUpdatableSystem
}
