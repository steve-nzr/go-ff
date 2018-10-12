package net

import (
	"fmt"
	"net"
	"time"
)

// Server represents a TCP server
type Server struct {
	connConnected    ConnectionActionHandler
	connDisconnected ConnectionActionHandler
	connMessage      ConnectionMessageHandler
	host             string
}

// ConnectionActionHandler for connection & disconnection
type ConnectionActionHandler func(c *Client)

// ConnectionMessageHandler for message
type ConnectionMessageHandler func(c *Client, p *Packet)

// Create a new Server
func Create(host string) *Server {
	server := new(Server)
	server.host = host

	return server
}

// OnConnected attach the specified function to the OnConnected event
func (ns *Server) OnConnected(fnc ConnectionActionHandler) *Server {
	ns.connConnected = fnc
	return ns
}

// OnDisconnected attach the specified function to the OnDisconnected event
func (ns *Server) OnDisconnected(fnc ConnectionActionHandler) *Server {
	ns.connDisconnected = fnc
	return ns
}

// OnMessage attach the specified function to the OnMessage event
func (ns *Server) OnMessage(fnc ConnectionMessageHandler) *Server {
	ns.connMessage = fnc
	return ns
}

// Start the server
func (ns *Server) Start() {
	l, err := net.Listen("tcp", ns.host)
	if err != nil {
		panic(err.Error)
	}

	defer l.Close()

	fmt.Println("Listening on " + ns.host)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}

		go handleClient(ns, conn)
	}
}

func handleClient(ns *Server, c net.Conn) {
	netClient := newClient(c)
	ns.connConnected(netClient)

	defer netClient.Close()
	defer ns.connDisconnected(netClient)

	for {
		buf := make([]byte, 2048)
		readLength, err := netClient.Read(buf)
		if err != nil || readLength < 1 {
			fmt.Println("Error reading:", err.Error())
			break
		}

		go handleMessage(ns, netClient, uint32(readLength), buf)
		time.Sleep(5 * time.Millisecond)

		if netClient.Conn == nil {
			return
		}
	}
}

func handleMessage(ns *Server, c *Client, readLen uint32, buf []byte) {
	packet := ReadPacket(buf)
	if packet.ReadUInt8() != 0x5E {
		fmt.Println("Invalid header")
		return
	}

	// CHECKSUM SKIP
	packet.ReadUInt32()

	length := packet.ReadUInt32()
	if length != uint32(len(buf[:readLen]))-13 {
		fmt.Println("Invalid size")
		return
	}

	// CHECKSUM SKIP
	packet.ReadUInt32()

	ns.connMessage(c, packet)
}
