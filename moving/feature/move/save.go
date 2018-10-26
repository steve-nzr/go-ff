package move

import (
	"flyff/common/service/cache"
)

// SaveMovingComponent saves only movin_ fields
func SaveMovingComponent(p *cache.Player) {
	cache.Connection.Model(p).Where("net_client_id = ?", p.NetClientID).Updates(map[string]interface{}{
		"movin_x":      p.Moving.Destination.X,
		"movin_y":      p.Moving.Destination.Y,
		"movin_z":      p.Moving.Destination.Z,
		"movin_motion": p.Moving.Motion,
		"movin_angle":  p.Moving.Angle,
	})
}

// SavePosition saves only posit_ fields
func SavePosition(p *cache.Player) {
	cache.Connection.Model(p).Where("net_client_id = ?", p.NetClientID).Updates(map[string]interface{}{
		"posit_x": p.Position.Vec.X,
		"posit_y": p.Position.Vec.Y,
		"posit_z": p.Position.Vec.Z,
	})
}
