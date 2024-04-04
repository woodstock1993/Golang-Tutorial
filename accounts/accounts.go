package accounts

import (
	"errors"
	"fmt"
)


// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Cant withdraw")
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the Account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

//
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\n Has: ", a.Balance())
}
