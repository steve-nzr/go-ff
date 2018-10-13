package in

import (
	"flyff/core/net"
	"flyff/world/entities"
	"flyff/world/packets/format"
	"flyff/world/packets/out"
	"flyff/world/service/gamemap"
)

func Chat(pe *entities.PlayerEntity, p *net.Packet) {
	var chat format.Chat
	chat.Construct(p)

	chatPacket := out.MakeChat(pe, &chat)

	gamemap.Manager.SendFrom(pe, &chatPacket)
	pe.Send(chatPacket)
}
