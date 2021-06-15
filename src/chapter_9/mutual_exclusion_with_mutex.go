package bank3

import "sync"

type guardedBalance struct {
	mutex  sync.Mutex
	amount int
}

func newBalance(initialBalance int) guardedBalance {
	return guardedBalance{
		mutex:  sync.Mutex{},
		amount: 0,
	}
}

var balance = newBalance(0)

func Deposit(amount int) {
	balance.mutex.Lock()
	// RAII would be a better solution than defer in this case.
	// Example: https://rust-unofficial.github.io/patterns/patterns/behavioural/RAII.html
	defer balance.mutex.Unlock()

	balance.amount += amount
}

func Balance() int {
	balance.mutex.Lock()
	defer balance.mutex.Unlock()

	return balance.amount
}
