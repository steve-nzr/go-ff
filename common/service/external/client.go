package external

import (
	"go-ff/common/def/packet/packettype"
	"log"
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
func (nc *Client) Send(p *Packet) {
	p.Finalize()

	_, err := nc.Write(p.Data[:p.Offset])
	if err != nil {
		log.Println(err)
		nc.Close()
	}
}

// SendFinalized sends a Packet to the Client
func (nc *Client) SendFinalized(p *Packet) {
	_, err := nc.Write(p.Data[:p.Offset])
	if err != nil {
		log.Println(err)
		nc.Close()
	}
}

// GenerateID "pseudo-random" from time
func GenerateID() uint32 {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	return rnd.Uint32()
}

// SendGreetings to the client
func (nc *Client) SendGreetings() *Client {
	p := MakePacket(packettype.Welcome).
		WriteUInt32(nc.ID)

	nc.Send(p)
	return nc
}
