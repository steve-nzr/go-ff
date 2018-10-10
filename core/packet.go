package core

import (
	"encoding/binary"
	"math"
	"strings"
)

// Packet stores data & offset for a in/out message
type Packet struct {
	data             []byte
	offset           uint32
	mergePacketCount uint16
}

const (
	// GREETINGS is the Welcome packet
	GREETINGS OutPacketType = 0x00
	// SERVERLIST contains all informations to list servers & channels
	SERVERLIST OutPacketType = 0xFD
	// WORLDADDR contains the address of the selected worldserver
	WORLDADDR OutPacketType = 0xF2
	// PLAYERLIST contains all informations to list servers & channels
	PLAYERLIST OutPacketType = 0xF3
	// PREJOIN is the authorization to connect to WorldServer
	PREJOIN OutPacketType = 0xFF05
)

const (
	ENVIRONMENTALL SnapshotPacket = 0x0063
	WORLDREADINFO  SnapshotPacket = 0x9910
	ADDOBJ         SnapshotPacket = 0x00F0
)

// OutPacketType is an outgoing type of packet
type OutPacketType uint32

// SnapshotPacket is an outgoing type of packet from world
type SnapshotPacket uint16

// StartMergePacket create a simple packet without protocol & data
func StartMergePacket(moverID uint32, cmd uint16, mainCmd uint32) Packet {
	packet := Packet{make([]byte, 4096), 0, 1}
	return packet.
		WriteUInt8(0x5E).
		WriteUInt32(0).
		WriteUInt32(mainCmd).
		WriteUInt32(0).
		WriteUInt16(packet.mergePacketCount). // MergePacketcount
		WriteUInt32(moverID).
		WriteUInt16(uint16(cmd))
}

// MakePacket constructs a new packet with the given protocol
func MakePacket(protocol OutPacketType) Packet {
	p := Packet{make([]byte, 4096), 0, 0}.
		WriteUInt8(0x5E).
		WriteUInt32(0).
		WriteUInt32(uint32(protocol))

	return p
}

// AddMergePart add a new entry to the packet
func (p Packet) AddMergePart(protocol SnapshotPacket, moverID uint32) Packet {
	p.mergePacketCount++

	lastOffset := p.offset
	p.offset = 13
	p = p.WriteUInt16(p.mergePacketCount)
	p.offset = lastOffset

	p = p.WriteUInt32(moverID).
		WriteUInt16(uint16(protocol))

	return p
}

// Finalize prepares the packet to be send by adding size
func (p Packet) finalize() Packet {
	totalLen := p.offset
	if totalLen < 5 {
		return p
	}

	binary.LittleEndian.PutUint32(p.data[1:], uint32(totalLen-5))
	return p
}

// WriteFloat32 at the current offset
func (p Packet) WriteFloat32(i float32) Packet {
	binary.LittleEndian.PutUint32(p.data[p.offset:], math.Float32bits(i))
	p.offset += (32 / 8)
	return p
}

// WriteInt64 at the current offset
func (p Packet) WriteInt64(i int64) Packet {
	binary.LittleEndian.PutUint64(p.data[p.offset:], uint64(i))
	p.offset += (64 / 8)
	return p
}

// WriteUInt64 at the current offset
func (p Packet) WriteUInt64(i uint64) Packet {
	binary.LittleEndian.PutUint64(p.data[p.offset:], i)
	p.offset += (64 / 8)
	return p
}

// WriteInt32 at the current offset
func (p Packet) WriteInt32(i int32) Packet {
	binary.LittleEndian.PutUint32(p.data[p.offset:], uint32(i))
	p.offset += (32 / 8)
	return p
}

// WriteUInt32 at the current offset
func (p Packet) WriteUInt32(i uint32) Packet {
	binary.LittleEndian.PutUint32(p.data[p.offset:], i)
	p.offset += (32 / 8)
	return p
}

// WriteInt16 at the current offset
func (p Packet) WriteInt16(i int16) Packet {
	binary.LittleEndian.PutUint16(p.data[p.offset:], uint16(i))
	p.offset += (16 / 8)
	return p
}

// WriteUInt16 at the current offset
func (p Packet) WriteUInt16(i uint16) Packet {
	binary.LittleEndian.PutUint16(p.data[p.offset:], i)
	p.offset += (16 / 8)
	return p
}

// WriteInt8 at the current offset
func (p Packet) WriteInt8(i int8) Packet {
	p.data[p.offset] = uint8(i)
	p.offset += (8 / 8)
	return p
}

// WriteUInt8 at the current offset
func (p Packet) WriteUInt8(i uint8) Packet {
	p.data[p.offset] = i
	p.offset += (8 / 8)
	return p
}

// WriteString (size+string) at the current offset
func (p Packet) WriteString(s string) Packet {
	length := len(s)
	if length < 1 {
		return p
	}

	p = p.WriteInt32(int32(length))
	for i := 0; i < length; i++ {
		p = p.WriteUInt8(s[i])
	}

	return p
}

// ReadPacket create a new packet instance with the given input data
func ReadPacket(d []byte) *Packet {
	p := new(Packet)
	p.data = d

	return p
}

// ReadInt32 at the current offset
func (p *Packet) ReadInt32() int32 {
	i := binary.LittleEndian.Uint32(p.data[p.offset:])
	p.offset += (32 / 8)
	return int32(i)
}

// ReadUInt32 at the current offset
func (p *Packet) ReadUInt32() uint32 {
	i := binary.LittleEndian.Uint32(p.data[p.offset:])
	p.offset += (32 / 8)
	return i
}

// ReadInt16 at the current offset
func (p *Packet) ReadInt16() int16 {
	i := binary.LittleEndian.Uint16(p.data[p.offset:])
	p.offset += (16 / 8)
	return int16(i)
}

// ReadUInt16 at the current offset
func (p *Packet) ReadUInt16() uint16 {
	i := binary.LittleEndian.Uint16(p.data[p.offset:])
	p.offset += (16 / 8)
	return i
}

// ReadInt8 at the current offset
func (p *Packet) ReadInt8() int8 {
	i := p.data[p.offset]
	p.offset += (8 / 8)
	return int8(i)
}

// ReadUInt8 at the current offset
func (p *Packet) ReadUInt8() uint8 {
	i := p.data[p.offset]
	p.offset += (8 / 8)
	return i
}

// ReadString (size+string) at the current offset
func (p *Packet) ReadString() string {
	var buffer strings.Builder
	len := p.ReadUInt32()

	for i := uint32(0); i < len; i++ {
		buffer.WriteByte(p.ReadUInt8())
	}

	return buffer.String()
}
