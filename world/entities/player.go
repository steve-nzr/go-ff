package entities

import (
	"flyff/core/net"
	mc "flyff/world/feature/movement/component"
	"flyff/world/game/component"
)

type PlayerEntity struct {
	Entity
	Client     *net.Client
	PlayerID   uint32
	Slot       uint8
	JobID      uint8
	HairColor  uint32
	HairID     uint32
	SkinSetID  uint32
	FaceID     uint32
	Statistics component.Statistics
	Moving     mc.Moving
}
