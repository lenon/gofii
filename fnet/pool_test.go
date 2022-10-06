package fnet_test

import (
	"errors"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/lenon/gofii/fnet"
	"github.com/stretchr/testify/assert"
)

func produceString(item int) (string, error) {
	// sleep with a random interval to emulate a real-world scenario
	time.Sleep(time.Duration(100) * time.Millisecond)

	return fmt.Sprintf("item %d", item), nil
}

func produceStringWithError(item int) (string, error) {
	if item == 1 {
		return "", errors.New("something went terribly wrong")
	}

	return fmt.Sprintf("item %d", item), nil
}

func consumeAll(jobs int, out <-chan fnet.JobResult[string]) []fnet.JobResult[string] {
	results := []fnet.JobResult[string]{}

	for i := 1; i <= jobs; i++ {
		jobResult := <-out
		results = append(results, jobResult)
	}

	// sort results to avoid flaky tests
	sort.Slice(results, func(i, j int) bool {
		return results[i].Result < results[j].Result
	})

	return results
}

func TestStartPool(t *testing.T) {
	concurrency := 2
	jobs := 5
	in, out := fnet.StartPool(concurrency, jobs, produceString)

	for i := 1; i <= jobs; i++ {
		in <- i
	}
	close(in)

	results := consumeAll(jobs, out)

	expected := []fnet.JobResult[string]{
		{Result: "item 1", Err: nil},
		{Result: "item 2", Err: nil},
		{Result: "item 3", Err: nil},
		{Result: "item 4", Err: nil},
		{Result: "item 5", Err: nil},
	}

	assert.Equal(t, expected, results)
}

func TestStartPoolWithError(t *testing.T) {
	concurrency := 2
	jobs := 2
	in, out := fnet.StartPool(concurrency, jobs, produceStringWithError)

	for i := 1; i <= jobs; i++ {
		in <- i
	}
	close(in)

	results := consumeAll(jobs, out)

	expected := []fnet.JobResult[string]{
		{Result: "", Err: errors.New("something went terribly wrong")},
		{Result: "item 2", Err: nil},
	}

	assert.Equal(t, expected, results)
}
