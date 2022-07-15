package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/mojochao/go-workerpool/pkg/workerpool"
)

// worker is a workerpool worker function processing jobs, here a channel
// of int inputs, and results, here a channel of int outputs. where each
// output value is double that of each input value.
func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(time.Second) // doubling ints is hard work
		results <- job * 2
	}
}

func main() {
	// Create a workerpool configured with the number of workers to pool.
	wp := workerpool.New(5, worker)

	// Define jobs, here a slice of ints to be doubled.
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Run the jobs in the workerpool workers and print the doubled results.
	results := wp.Run(jobs)
	fmt.Println(results)
	sort.Ints(results)
	fmt.Println(results)
}
