package main

import (
	"bank/pkg/bank/payment"
	"bank/pkg/bank/types"
)

func main() {
	//card := types.Card{
	//	Id:       12345,
	//	PAN:      "5058 xxxx xxxx 9999",
	//	Balance:  999_99,
	//	Currency: "TJS",
	//	Color:    "white",
	//	Name:     "Infinity",
	//	Active:   true,
	//}
	//
	//id := card.Id
	//fmt.Println(id)
	//card.Balance = 2000_00

	//fmt.Println(card)
	items := []types.Payment{
		{ID: 1,
			Amount: 10_000_00},
		{ID: 2,
			Amount: 5_000_00},
		{ID: 3,
			Amount: 11_000_00},
		{ID: 3,
			Amount: 0},
	}
	payment.Max(items)
}
