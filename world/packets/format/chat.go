package format

import "flyff/core/net"

type Chat struct {
	Message string
}

func (c *Chat) Construct(p *net.Packet) {
	c.Message = p.ReadString()
}
