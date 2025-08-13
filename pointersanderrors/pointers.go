package pointersanderrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance	Bitcoin
}

var ErrInsufficientFund = "cannot withdraw, insufficient balance"

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC",b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("The memory address of the wallet balance variable - %p\n", &w.balance)
	(*w).balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New(ErrInsufficientFund)
	}
	w.balance -= amount
	return nil
}