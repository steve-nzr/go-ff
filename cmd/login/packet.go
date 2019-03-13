package main

import (
	"go-ff/pkg/def/packet/packettype"
	"math"

	"go-ff/pkg/service/external"
)

func sendServerList(c *external.Client) {
	p := external.MakePacket(packettype.Srvr_list).
		WriteUInt32(0).
		WriteUInt8(1).
		WriteString("test").
		WriteUInt32(2)

	for i, server := range servers {
		p.WriteUInt32(math.MaxUint32).
			WriteInt32(int32(i + 1)).
			WriteString(server.name).
			WriteString(server.ip).
			WriteUInt32(0).
			WriteUInt32(0).
			WriteUInt32(1).
			WriteUInt32(0)

		for j, channel := range server.channels {
			p.WriteUInt32(uint32(i + 1)).
				WriteUInt32(uint32(j + 1)).
				WriteString(channel.name).
				WriteString(channel.ip).
				WriteUInt32(0).
				WriteUInt32(0).
				WriteUInt32(1).
				WriteUInt32(channel.maxPlayer)
		}
	}

	c.Send(p)
}
