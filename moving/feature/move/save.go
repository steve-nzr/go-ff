package move

import (
	"flyff/common/service/cache"
)

// SaveMovingPosition saves only movin_ fields
func SaveMovingPosition(p *cache.Player) {
	cache.Connection.Model(p).Where("net_client_id = ?", p.NetClientID).Updates(map[string]interface{}{
		"movin_x": p.Moving.Vec.X,
		"movin_y": p.Moving.Vec.Y,
		"movin_z": p.Moving.Vec.Z,
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
