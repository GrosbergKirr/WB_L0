package cache

import (
	"project0/models"
)

// Make cache
type MakeCacher interface {
	GetAllFromDb() []models.Order
}

//type Cache struct {
//	mp map[string]models.Order
//}

func MakeCache(mk MakeCacher, cache map[string]models.Order) map[string]models.Order {
	orderarray := mk.GetAllFromDb()

	for i := range orderarray {
		cache[orderarray[i].OrderUid] = orderarray[i]
	}

	return cache

}
