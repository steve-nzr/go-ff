package entities

import (
	"encoding/json"
	"flyff/core/net"
	mc "flyff/world/feature/movement/component"
	"flyff/world/game/component"
	"flyff/world/service/messaging/channel"
	"flyff/world/service/messaging/definitions"

	"github.com/streadway/amqp"
)

type PlayerEntity struct {
	Entity
	NetClientID uint32 `gorm:"primary_key"`
	PlayerID    uint32
	Slot        uint8
	JobID       uint8
	HairColor   uint32
	HairID      uint32
	SkinSetID   uint32
	FaceID      uint32
	Statistics  component.Statistics
	Moving      mc.Moving
}

func (pe *PlayerEntity) Send(p net.Packet) {
	var po definitions.ExternalPacket
	po.Packet = p.Finalize()
	po.To = []uint32{pe.NetClientID}
	bytes, _ := json.Marshal(&po)

	channel.Channel.Publish(
		"packet_out",
		"#",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		},
	)
}
