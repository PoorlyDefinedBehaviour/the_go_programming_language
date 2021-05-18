package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

/*
The defer statement is usually used for cleanup.

Deferred functions are run in reverse order.

Deferred functions are run before the stack is unwound.



example:

package ioutil

func ReadFile(filename string) ([]byte, error){
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ReadAll(file)
}
*/

func title(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	/*
		Deferred functions are called after the function that contains the defer
		statement has finished, whether normally, by executing a return statement
		or falling off the end, or abnormally, by panicking.
	*/
	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return fmt.Errorf("%s has content-type %s, not text/html", url, contentType)
	}

	document, err := html.Parse(response.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// do something with the document
	fmt.Println(document)

	return nil
}

var (
	mutex sync.Mutex
	m     = make(map[string]int)
)

// using defer to always unlock a mutex
func lookup(key string) int {
	mutex.Lock()
	defer mutex.Unlock()
	return m[key]
}

func trace(message string) func() {
	start := time.Now()

	log.Printf("enter %s", message)

	return func() {
		log.Printf("exit %s (%s)", message, time.Since(start))
	}
}

func f() {
	defer trace("f()")()
}

/*
Deferred functions run after return statements have updated the function's
result variables. Because an anonymous function can access its
enclosing function's variables, including named results,
a deferred anonymouws function can observe the function's results.
*/
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

/*
Deferred anonymous functions can change the values that the enclosing
function returns to its caller:
*/
func triple(x int) (result int) {
	defer func() {
		result += x
	}()

	return double(x)
}

/*
Pitfalls:
Using defer inside a loop
*/
func deferInsideLoopBad(filenames []string) error {
	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}

		// Will only be called when the function finishes.
		// Since the file is not closed after each iteration
		// we may run out of file descriptors.
		defer file.Close()

		// do something with the file
	}

	return nil
}

/*
If you need to defer after each loop iteration,
create a function and move the defer statement inside of it.
*/
func deferInsideLoopGood(filenames []string) error {
	for _, filename := range filenames {
		err := doFile(filename)
		if err != nil {
			return err
		}
	}

	return nil
}

func doFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	// Will only be called when the function finishes.
	// Since the file is not closed after each iteration
	// we may run out of file descriptors.
	defer file.Close()

	// do something with the file

	return nil
}

func main() {
	f()

	double(2)

	fmt.Println(triple(2))
}
