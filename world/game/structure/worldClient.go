package structure

import "flyff/core"

type WorldClient struct {
	*core.NetClient
	PlayerEntity *PlayerEntity
}
