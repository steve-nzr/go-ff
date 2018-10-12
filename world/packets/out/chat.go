package out

import (
	"flyff/core/net"
	"flyff/world/game/structure"
	"flyff/world/packets/format"
)

func MakeChat(wc *structure.WorldClient, chat *format.Chat) net.Packet {
	return net.StartMergePacket(uint32(wc.PlayerEntity.ID), uint16(0x0001), 0xFFFFFF00).
		WriteString(chat.Message)
}
