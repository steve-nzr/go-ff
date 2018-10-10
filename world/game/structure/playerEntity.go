package structure

import "flyff/world/game/component"

type PlayerEntity struct {
	Entity
	Slot       uint8
	JobID      uint8
	HairColor  uint32
	HairID     uint32
	SkinSetID  uint32
	FaceID     uint32
	Statistics component.Statistics
}
