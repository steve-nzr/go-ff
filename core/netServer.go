package core

import (
	"fmt"
	"net"
	"time"
)

// ConnectionSingleHandler when new connection initiated
type ConnectionSingleHandler func(c *NetClient)

// ConnectionMessageHandler when new message arrives
type ConnectionMessageHandler func(c *NetClient, p *Packet)

// NetServerConfig holds configuration to run a NetServer
type NetServerConfig struct {
	Host                string
	Port                string
	Type                NetServerType
	ConnectionClosed    ConnectionSingleHandler
	ConnectionInitiated ConnectionSingleHandler
	ConnectionMessage   ConnectionMessageHandler
}

const (
	// NetServerTCP for TCP
	NetServerTCP NetServerType = "tcp"
	// NetServerUDP for UDP
	NetServerUDP NetServerType = "udp"
)

// NetServerType is the NetServer socket type
type NetServerType string

var netServerConfig NetServerConfig

// StartNetServer start listening with the given configuration
// It also calls onConnected function at each new connection
func StartNetServer(nc NetServerConfig) {
	netServerConfig = nc

	l, err := net.Listen(string(netServerConfig.Type), netServerConfig.Host+":"+netServerConfig.Port)
	if err != nil {
		panic("Error listening: " + err.Error())
	}
	defer l.Close()

	fmt.Println("Listening on " + netServerConfig.Host + ":" + netServerConfig.Port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}

		go connectionInitiated(conn)
	}
}

func connectionMessage(c *NetClient, readLen uint32, buf []byte) {
	//startTime := time.Now()

	packet := ReadPacket(buf)
	header := packet.ReadUInt8()
	if header != 0x5E {
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

	netServerConfig.ConnectionMessage(c, packet)

	//fmt.Printf("Packet processed in : %fms\n", float64(time.Now().Sub(startTime).Nanoseconds())/1000000.0)
}

func connectionInitiated(c net.Conn) {
	netClient := newNetClient(c)
	defer netClient.Close()
	defer netServerConfig.ConnectionClosed(netClient)

	netServerConfig.ConnectionInitiated(netClient)

	for {
		buf := make([]byte, 2048)

		readLength, err := netClient.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		if readLength < 1 {
			fmt.Println("Error reading:", err.Error())
			break
		}

		go connectionMessage(netClient, uint32(readLength), buf)

		// Closed ?
		if netClient.Conn == nil {
			return
		}

		time.Sleep(5 * time.Millisecond)
	}
}
