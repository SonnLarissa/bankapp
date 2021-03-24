package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 10,
		},
		{
			ID: 2,
			Amount: 0,
		},
		{
			ID: 99,
			Amount: 99,
		},
	}

	max := Max(payments)
	fmt.Println(max.ID)
	//Output:99
}
