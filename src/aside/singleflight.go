package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

/*
The singleflight package provides a duplicate function call
suppression mechanism, which means if multiple goroutines call
the same function concurrently, every function returns
the same return returned by the first caller.
*/

func task() (interface{}, error) {
	fmt.Println("task()")
	time.Sleep(time.Second * 5)

	return "done", nil
}

func main() {
	singleflightGroup := singleflight.Group{}

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer waitGroup.Done()

			// task() is only called once
			value, _, _ := singleflightGroup.Do("task", task)

			// prints "done" five times because it is the value returned
			// by task()
			fmt.Println(value.(string))
		}()
	}

	waitGroup.Wait()
}
