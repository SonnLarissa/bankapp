package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	//Output: 1000000
}
func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	//Output: 0
}
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	//Output:2000000
}
func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 20_000_01)
	fmt.Println(result.Balance)
	//Output: 2000000
}

func ExampleDeposit_positive() {
	card := &types.Card{Balance: 20_000_00, Active: true}
	Deposit(card, 5000_00)
	fmt.Println(card.Balance)
	//Output: 2500000
}
func ExampleDeposit_inactive() {
	card := &types.Card{Balance: 20_000_00, Active: false}
	Deposit(card, 5000_00)
	fmt.Println(card.Balance)
	//Output: 2000000
}
func ExampleDeposit_limit() {
	card := &types.Card{Balance: 20_000_00, Active: false}
	Deposit(card, 50_000_10)
	fmt.Println(card.Balance)
	//Output: 2000000
}
func ExampleAddBonus_positive() {
	card := &types.Card{Balance: 20_000_00, MinBalance: 20_000_00, Active: true}
	AddBonus(card, 3, 31, 365)
	fmt.Println(card.Balance)
	//Output: 2005095
}
func ExampleAddBonus_inactive() {
	card := &types.Card{Balance: 20_000_00, Active: false}
	AddBonus(card, 3, 31, 365)
	fmt.Println(card.Balance)
	//Output: 2000000
}
func ExampleAddBonus_negativeBalance() {
	card := &types.Card{Balance: -1, MinBalance: -1, Active: true}
	AddBonus(card, 3, 31, 365)
	fmt.Println(card.Balance)
	//Output: -1
}
func ExampleAddBonus_limit() {
	card := &types.Card{Balance: 100_000_000, MinBalance: 100_000_000, Active: true}
	AddBonus(card, 3, 31, 365)
	fmt.Println(card.Balance)
	//Output: 100254794
}

func ExampleAddBonus_limitSalom() {
	card := &types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: 1002465
}
func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true},
		{
			Balance: 10_000_00,
			Active:  false,
		},
	}
	fmt.Println(Total(cards))
	//Output:2000000
}
func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
			PAN:     "salom",
		},
		{
			Balance: 2_000_00,
			Active:  true,
			PAN:     "depozit",
		},
		{
			Balance: 1_000_00,
			Active:  false,
			PAN:     "alek",
		},
		{
			Balance: -1_000_00,
			Active:  true,
			PAN:     "qwerty",
		},
	}
	r := PaymentSources(cards)
	for _, source := range r {
		fmt.Println(source.Number)
	}
	//Output:
	//salom
	//depozit
	//
	//
}
