package out

import (
	"flyff/core"
	"flyff/world/game/structure"
	"flyff/world/packets/format"
)

func MakeChat(wc *structure.WorldClient, chat *format.Chat) core.Packet {
	return core.StartMergePacket(uint32(wc.Character.ID), uint16(0x0001), 0xFFFFFF00).
		WriteString(chat.Message)
}
