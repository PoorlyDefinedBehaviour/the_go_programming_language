package main

var (
	// buffered channel with capacity of 1:
	// Sender blocks when len(channel) == 1.
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	// blocks until whoever is accessing `balance` is done with it.
	// mutex comparisson:
	// blocks until lock is acquired.
	sema <- struct{}{}

	balance += amount

	// release lock
	<-sema
}

func main() {

}
