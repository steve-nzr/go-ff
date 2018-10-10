package format

import "flyff/core"

type Chat struct {
	Message string
}

func (c *Chat) Construct(p *core.Packet) {
	c.Message = p.ReadString()
}
