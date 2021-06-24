package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

type JobQueueSemaphore struct {
	jobs      chan func()
	semaphore *semaphore.Weighted
}

func NewJobQueueSemaphore(concurrency int64) JobQueueSemaphore {
	jobQueue := JobQueueSemaphore{
		jobs:      make(chan func(), concurrency),
		semaphore: semaphore.NewWeighted(concurrency),
	}

	go jobQueue.executeJobs()

	return jobQueue
}

func (jobQueue *JobQueueSemaphore) executeJobs() {
	for job := range jobQueue.jobs {
		err := jobQueue.semaphore.Acquire(context.Background(), 1)
		if err != nil {
			fmt.Println(err)
		}

		go func() {
			job()
			jobQueue.semaphore.Release(1)
		}()
	}
}

func (jobQueue *JobQueueSemaphore) Add(job func()) {
	jobQueue.jobs <- job
}

func main() {
	jobQueue := NewJobQueueSemaphore(3)

	for i := 0; i < 9; i++ {
		jobQueue.Add(func() {
			time.Sleep(1 * time.Second)
			fmt.Println("done executing")
		})
	}

	time.Sleep(10 * time.Second)
}
