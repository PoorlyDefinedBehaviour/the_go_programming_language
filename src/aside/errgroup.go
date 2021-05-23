package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

/*
errgroup provides synchronization, error propagation, and Context
cancelation for groups of goroutines working on subtasks.
*/

func requestAPI(i int) error {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i)

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	_, err = io.Copy(ioutil.Discard, response.Body)

	return err
}

func withoutErrGroup() {
	tasks := 10
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(tasks)

	for i := 0; i < tasks; i++ {
		go func(i int) {
			defer waitGroup.Done()

			if err := requestAPI(i); err != nil {
				print(err.Error())
			}
		}(i)
	}

	waitGroup.Wait()
}

func withErrGroup() {
	group, _ := errgroup.WithContext(context.TODO())

	for i := 0; i < 10; i++ {
		currentIndex := i

		group.Go(func() error {
			return requestAPI(currentIndex)
		})
	}

	if err := group.Wait(); err != nil {
		print(err.Error())
	}
}

func main() {
}
