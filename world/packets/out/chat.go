package out

import (
	"flyff/core/net"
	"flyff/world/entities"
	"flyff/world/packets/format"
)

func MakeChat(pe *entities.PlayerEntity, chat *format.Chat) net.Packet {
	return net.StartMergePacket(uint32(pe.ID), uint16(0x0001), 0xFFFFFF00).
		WriteString(chat.Message)
}
