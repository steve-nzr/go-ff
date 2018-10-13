package entities

import (
	"flyff/world/feature/movement/component"
)

type Entity struct {
	ID       uint32
	Type     uint8
	Name     string
	Gender   uint8
	ModelID  uint32
	Position component.Position
	Angle    float32
	Size     uint8
	Level    uint16
}
