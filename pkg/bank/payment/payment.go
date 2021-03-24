package payment

import (
	"bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment {
	var maxItem types.Payment
	index := 0
	for i := 0; i < len(payments); i++ {
		if maxItem.Amount < payments[i].Amount {
			maxItem = payments[i]
			index = i
		}
	}
	return payments[index]
}
