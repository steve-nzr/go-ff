package core

// PacketStruct represents an incoming packet interface with deserialization function
type PacketStruct interface {
	construct(p Packet) PacketStruct
}
