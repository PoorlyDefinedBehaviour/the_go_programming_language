package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

var functionTimedOutErr = errors.New("function timed out")

func withTimeout(timeout time.Duration, f func() interface{}) (interface{}, error) {
	returnValueChan := make(chan interface{}, 1)

	go func() {
		returnValueChan <- f()
	}()

	var returnValue interface{}

	select {
	case value := <-returnValueChan:
		return value, nil
	case <-time.After(timeout):
		return returnValue, functionTimedOutErr
	}
}

func main() {
	value, err := withTimeout(1*time.Second, func() interface{} {
		time.Sleep(3 * time.Second)
		return 1
	})

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(value)
}
