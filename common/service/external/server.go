package external

import (
	"log"
	"net"
	"time"
)

// PacketHandler represents an incoming packet from a client
type PacketHandler struct {
	ClientID uint32
	Packet   *Packet
}

// PacketEmitter represents an outgoing packet to the ConnectionService
type PacketEmitter struct {
	To     []uint32
	Packet *Packet
}

// Server represents a TCP server
type Server struct {
	connConnected    chan<- *Client
	connDisconnected chan<- *Client
	connMessage      chan<- *PacketHandler
	host             string
}

// Create a new Server
func Create(host string) *Server {
	server := new(Server)
	server.host = host

	return server
}

// OnConnected attach the specified function to the OnConnected event
func (ns *Server) OnConnected(c chan<- *Client) *Server {
	ns.connConnected = c
	return ns
}

// OnDisconnected attach the specified function to the OnDisconnected event
func (ns *Server) OnDisconnected(c chan<- *Client) *Server {
	ns.connDisconnected = c
	return ns
}

// OnMessage attach the specified function to the OnMessage event
func (ns *Server) OnMessage(c chan<- *PacketHandler) *Server {
	ns.connMessage = c
	return ns
}

// Start the server
func (ns *Server) Start() {
	l, err := net.Listen("tcp", ns.host)
	if err != nil {
		panic(err.Error)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			continue
		}

		go ns.handleClient(conn)
	}
}

func (ns *Server) handleClient(c net.Conn) {
	netClient := newClient(c)
	ns.connConnected <- netClient

	defer func() {
		ns.connDisconnected <- netClient
		netClient.Close()
	}()

	for {
		buf := make([]byte, 2048)
		readLength, err := netClient.Read(buf)
		if err != nil || readLength < 1 {
			break
		}

		go func() {
			var i uint32

			for {
				if i >= uint32(readLength) {
					break
				}

				packet := ReadPacket(buf[i:])
				if packet.ReadUInt8() != 0x5E {
					break
				}

				// CHECKSUM SKIP
				packet.ReadUInt32()

				packetLen := packet.ReadUInt32()
				i += packetLen + 13
				nextBuf := buf[i:]
				if i < uint32(readLength) && len(nextBuf) > 0 && nextBuf[0] != 0x5E {
					break
				}

				// CHECKSUM SKIP
				packet.ReadUInt32()
				packetHandler := new(PacketHandler)
				packetHandler.ClientID = netClient.ID
				packetHandler.Packet = packet
				ns.connMessage <- packetHandler
			}
		}()

		time.Sleep(5 * time.Millisecond)

		if netClient.Conn == nil {
			return
		}
	}
}
