package main

import (
	"fmt"
	"time"
)

/*
In Go, each concurrently executing activity is called a goroutine.

For simplicity, we can assume that a goroutine is similar to a thread.

When a program starts, a goroutine calls the main function,
it is called the main goroutine.

New goroutines are created by the go statement. Syntactically,
a go statement is an ordinary function or method call prefixed
by the keyword go. A go statement causes the function to be
called in a newly created goroutine. The go statement itself
completes immediately.

f() // call f and wait for it to return
go f() // create a new goroutine the calls f() and don't wait for it to return
*/

func spinner(delay time.Duration) {
	for {
		for _, character := range `-\|/` {
			fmt.Printf("\r%c", character)
			time.Sleep(delay)
		}
	}
}

func slowFib(n int) int {
	if n < 2 {
		return n
	}

	return slowFib(n-1) + slowFib(n-2)
}

/*
When the main function returns, all goroutines are
abruptly terminated and the program exits.

Other than exiting the program, there is no programmatic way
for one goroutine to stop another. But there are ways
to communicate with a goroutine to request that it stop itself.
*/
func main() {
	go spinner(100 * time.Millisecond)

	const n = 45

	// isn't this line racing with the fmt.Printf call inside spinner()
	// since both are printing to stdout?
	fmt.Printf("\rFibonacci(%d) = %d\n", n, slowFib(n))
}
