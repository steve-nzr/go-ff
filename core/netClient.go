package core

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// NetClient stores the network connection + the ID
type NetClient struct {
	net.Conn
	ID uint32
}

// NewNetClient returns a NetClient instance with the given net.Conn
func newNetClient(c net.Conn) *NetClient {
	nc := new(NetClient)
	nc.Conn = c
	nc.ID = GenerateID()

	return nc
}

// Send a Packet to the NetClient
func (nc *NetClient) Send(p Packet) {
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

func (nc *NetClient) SendGreetings() *NetClient {
	p := MakePacket(GREETINGS).
		WriteUInt32(nc.ID)

	nc.Send(p)
	return nc
}
