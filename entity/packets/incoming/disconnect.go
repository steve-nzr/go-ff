package incoming

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
)

// Disconnect from World
func Disconnect(p *external.PacketHandler) {
	cache.Connection.Where("net_client_id = ?", p.ClientID).Delete(&cache.Player{})
}
