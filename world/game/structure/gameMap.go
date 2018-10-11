package structure

// GameMap represents a Map and it's components
type GameMap struct {
	ID      uint32
	Players map[uint32]*PlayerEntity
}
