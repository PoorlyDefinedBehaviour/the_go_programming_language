package main

import "fmt"

var (
	deposits = make(chan int)
	balances = make(chan int)
)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

// func teller() is what's called a monitor goroutine.
// It guards some variable tha otherwise would cause race
// conditions when accessed concurrently.
func teller() {
	balance := 0

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func main() {
	go teller()

	go func() {
		Deposit(10)
		Deposit(10)
	}()

	Deposit(10)

	fmt.Println(Balance())
}
