package model

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID      string          `json:"id"`
	Balance decimal.Decimal `json:"balance"`
}

func (a *Account) String() string {
	return fmt.Sprintf("ID: %s, Balance: %s", a.ID, a.Balance)
}

type Transfer struct {
	Sender   string          `json:"sender"`
	Receiver string          `json:"receiver"`
	Amount   decimal.Decimal `json:"amount"`
}

func (t *Transfer) String() string {
	return fmt.Sprintf("%s -----[ %s ]----> %s", t.Sender, t.Amount, t.Receiver)
}

type Id struct {
	Account string `json:"account"`
}
