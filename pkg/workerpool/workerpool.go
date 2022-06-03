// Package workerpool provides an implementation of the worker pool pattern for a
// configurable number of arbitrary, generic jobs. For details on this pattern,
// refer to https://gobyexample.com/worker-pools.
package workerpool

// WorkerPool is a generic worker pool representation whose workers process inputs of
// type J and return outputs of type R.
type WorkerPool[J any, R any] struct {
	Size    int
	jobs    chan J
	results chan R
}

// New returns a new WorkerPool for workers processing inputs of type J and returning
// outputs of type R.
func New[J any, R any](size int, worker func(id int, jobs <-chan J, results chan<- R)) *WorkerPool[J, R] {
	jobs := make(chan J, size)
	results := make(chan R, size)
	for i := 0; i < size; i++ {
		go worker(i, jobs, results)
	}
	return &WorkerPool[J, R]{
		Size:    size,
		jobs:    jobs,
		results: results,
	}
}

// Run processes jobs in workers processing inputs of type J and returning
// outputs of type R.
func (wp *WorkerPool[J, R]) Run(jobs []J) (results []R) {
	for _, job := range jobs {
		wp.jobs <- job
	}
	close(wp.jobs)

	num := len(jobs)
	for i := 0; i < num; i++ {
		results = append(results, <-wp.results)
	}
	return results
}
