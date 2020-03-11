package capital

import "time"

type State struct {
	id         uint
	Month      uint
	Year       uint
	Balance    float32
	TotalShare int
	createdAt  time.Time
	updatedAt  time.Time
}

func (cs *State) ShareValue() float32 {
	return cs.Balance / float32(cs.TotalShare)
}
