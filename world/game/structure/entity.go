package structure

import "flyff/world/game/component"

type IEntity interface{}

type Entity struct {
	WorldClient *WorldClient
	ID          uint32
	Type        uint8
	Name        string
	Gender      uint8
	ModelID     uint32
	Position    component.Position
	Angle       float32
	Size        uint8
	Level       uint16
}
