package account

import (
	"errors"
	"fmt"
)

// Account struct
// Public (start with upper case)
// Private (start with lower case)
type Account struct {
	owner 	string 
	balance int 
}

// error should start with 'err' eg.) errSomething
var errNoMoney = errors.New("Can't withdraw, you are poor")

// NewAccount creates Account
// This is like "constructor" pattern in Golang
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit ( a : receiver)
// convention : receiver name has to be first letter of struct
// Golang copy stuffs, >> (a Account) should be like (a *Account)
// to use original one, not a copied one
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance showing 
func (a Account) Balance() int {
	return a.balance
}

// Withdraw from Account
// Golang doesn't have try-catch-except
// This is how Golang control the Exception
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// return errors.New("Can't withdraw, you are poor")
		return errNoMoney
	}
	a.balance -=amount
	return nil
}

// ChangeOwner to newOwner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner show owner
func (a Account) Owner() string {
	return a.owner
}

// This is like __str__ in python
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
} 