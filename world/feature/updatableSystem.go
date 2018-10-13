package feature

import "flyff/world/entities"

type IUpdatableSystem interface {
	Update(time int64, pe *entities.PlayerEntity)
}
