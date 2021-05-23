package main

import (
	"context"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

/*
The semaphore packages provides a weighted semaphore implementation.
*/

func process() {
	sem := semaphore.NewWeighted(4)

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(10)

	go func() {
		for i := 0; i < 5; i++ {
			if err := sem.Acquire(context.TODO(), 1); err != nil {
				panic(err)
			}

			time.Sleep(time.Second * 1)

			sem.Release(1)

			waitGroup.Done()
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			if err := sem.Acquire(context.TODO(), 3); err != nil {
				panic(err)
			}

			time.Sleep(time.Second * 3)

			sem.Release(3)

			waitGroup.Done()
		}
	}()
}

func main() {

}
