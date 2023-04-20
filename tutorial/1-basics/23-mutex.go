package main

import (
	"fmt"
	"sync"
	"time"
)

// Using locks to protect data from being
// accessed by more than one user at a time
// Locks are another option when you don't
// have to pass dat

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		fmt.Println("Not enough money in account")
	} else {
		fmt.Printf("%d withdrawn | Balance: %d\n", v, a.balance)
		a.balance -= v
	}
}

func main() {
	var acct Account
	acct.balance = 100
	println("Balance:", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}
