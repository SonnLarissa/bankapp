package types


type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Money int64
type PAN string

type Card struct {
Id       int
PAN      PAN
Balance  Money
Currency Currency
Color    string
Name     string
Active   bool
MinBalance Money
}
// Payment provides information about the payment
type Payment struct{
	ID int
	Amount Money
}