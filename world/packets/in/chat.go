package in

import (
	"flyff/core"
	"flyff/world/game/structure"
	"flyff/world/packets/format"
	"flyff/world/packets/out"
	"flyff/world/service/gamemap"
)

func Chat(wc *structure.WorldClient, p *core.Packet) {
	var chat format.Chat
	chat.Construct(p)

	chatPacket := out.MakeChat(wc, &chat)

	gamemap.Manager.SendFrom(wc, &chatPacket)
	wc.Send(chatPacket)
}
