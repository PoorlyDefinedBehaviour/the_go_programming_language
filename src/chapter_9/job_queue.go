package main

import (
	"fmt"
	"time"
)

type JobQueue struct {
	jobs   chan func()
	tokens chan struct{}
}

func NewJobQueue(concurrency int) JobQueue {
	jobQueue := JobQueue{
		jobs:   make(chan func(), concurrency),
		tokens: make(chan struct{}, concurrency),
	}

	go jobQueue.executeJobs()

	return jobQueue
}

func (jobQueue *JobQueue) executeJobs() {
	for job := range jobQueue.jobs {
		jobQueue.tokens <- struct{}{}

		go func() {
			job()
			<-jobQueue.tokens
		}()
	}
}

func (jobQueue *JobQueue) Add(job func()) {
	jobQueue.jobs <- job
}

func main() {
	jobQueue := NewJobQueue(3)

	for i := 0; i < 9; i++ {
		jobQueue.Add(func() {
			time.Sleep(1 * time.Second)
			fmt.Println("done executing")
		})
	}

	time.Sleep(10 * time.Second)
}
