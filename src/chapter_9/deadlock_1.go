package deadlock1

import (
	"fmt"
	"sync"
)

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

func WithdrawWithDeadlock(amount int) error {
	// locks mutex
	balance.mutex.Lock()
	defer balance.mutex.Unlock()

	// tries to lock mutex but deadlocks since
	// the mutex is already locked at line 40
	Deposit(-amount)

	if Balance() < 0 {
		Deposit(amount)
		return fmt.Errorf("there's not enough balance to withdraw amount: %d", amount)
	}

	return nil
}

// just perform operations manually to ensure
// everything can be done with a single call to mutex.Lock()
func WithdrawWithoutDeadlock(amount int) error {
	balance.mutex.Lock()
	defer balance.mutex.Unlock()

	if balance.amount-amount < 0 {
		return fmt.Errorf("there's not enough balance to withdraw amount: %d", amount)
	}

	balance.amount -= amount

	return nil
}
