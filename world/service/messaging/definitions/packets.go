package definitions

import "flyff/core/net"

type ExternalPacket struct {
	Packet net.Packet
	To     []uint32
}

type Todo uint8

const (
	AddTodo    Todo = 1
	RemoveTodo Todo = 2
)

type InternalPacket struct {
	ID     uint32
	Packet net.Packet
	Todo   Todo
}
