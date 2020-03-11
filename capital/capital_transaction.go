package capital

import "time"

type Transaction struct {
	id        uint
	date      time.Time
	amount    float32
	fee       float32
	createdAt time.Time
	updatedAt time.Time
}

func (ct *Transaction) NetProfit() float32 {
	return ct.amount - ct.fee
}
