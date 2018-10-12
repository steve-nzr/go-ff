package net

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// Client stores the network connection + the ID
type Client struct {
	net.Conn
	ID uint32
}

// NewClient returns a Client instance with the given net.Conn
func newClient(c net.Conn) *Client {
	nc := new(Client)
	nc.Conn = c
	nc.ID = GenerateID()

	return nc
}

// Send a Packet to the Client
func (nc *Client) Send(p Packet) {
	p = p.finalize()

	_, err := nc.Write(p.data[:p.offset])
	if err != nil {
		fmt.Println(err)
		nc.Close()
	}
}

func GenerateID() uint32 {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	return rnd.Uint32()
}

func (nc *Client) SendGreetings() *Client {
	p := MakePacket(GREETINGS).
		WriteUInt32(nc.ID)

	nc.Send(p)
	return nc
}
