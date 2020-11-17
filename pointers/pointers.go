package pointers

import (
	"errors"
	"fmt"
)

// A Bitcoin type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// A Wallet to hold all the bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit some bitcoins
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance of your wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds error
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdrawl of bitcoints
func (w *Wallet) Withdrawl(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
