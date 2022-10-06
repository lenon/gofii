package fnet

import (
	"log"
)

type JobResult[T any] struct {
	Result T
	Err    error
}

func Worker[In any, Out any](id int, in <-chan In, out chan<- JobResult[Out], job func(In) (Out, error)) {
	log.Printf("worker %d started", id)

	// wait for a job to arrive
	for item := range in {
		result, err := job(item)

		out <- JobResult[Out]{
			Result: result,
			Err:    err,
		}
	}

	log.Printf("worker %d finished", id)
}

func StartPool[In any, Out any](concurrency, numberOfJobs int, job func(In) (Out, error)) (chan<- In, <-chan JobResult[Out]) {
	in := make(chan In, numberOfJobs)
	out := make(chan JobResult[Out], numberOfJobs)

	// spawn N workers to limit the number of concurrent jobs
	for id := 1; id <= concurrency; id++ {
		go Worker(id, in, out, job)
	}

	return in, out
}
