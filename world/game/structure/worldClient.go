package structure

import "flyff/core/net"

type WorldClient struct {
	*net.Client
	PlayerEntity *PlayerEntity
}
