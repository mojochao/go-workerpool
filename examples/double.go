package main

import (
	"fmt"
	"time"

	"github.com/mojochao/go-workerpool/pkg/workerpool"
)

func main() {
	double := func(id int, jobs <-chan int, results chan<- int) {
		for j := range jobs {
			fmt.Println("worker", id, "started  job", j)
			time.Sleep(time.Second)
			fmt.Println("worker", id, "finished job", j)
			results <- j * 2
		}
	}

	p := workerpool.New(5, double)
	results := p.Run([]int{1, 2, 3, 4, 5})
	fmt.Println(results)
}
