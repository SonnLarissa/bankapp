package card

import (
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		Id:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Active && card.Balance > 0 {
		if amount > 0 && amount < card.Balance && amount <= 20_000_00 {
			card.Balance = card.Balance - amount
		}
	}
	return card
}

const DepositLimit = 50_000_00
const IncomeLimit = 5_000_00

func Deposit(card *types.Card, amount types.Money) {
	if card.Active && amount <= DepositLimit && amount > 0 {
		card.Balance += amount
	}
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active || card.Balance <= 0 || percent <= 0 {
		return
	}
	income := ((card.MinBalance * types.Money(percent)) / 100) * types.Money(daysInMonth) / types.Money(daysInYear)
	if income <= IncomeLimit {
		card.Balance += income
	}
}

func Total(cards []types.Card) types.Money {
	var sum types.Money = 0
	for _, item := range cards {
		if item.Balance > 0 && item.Active {
			sum += item.Balance
		}
	}
	return sum
}
