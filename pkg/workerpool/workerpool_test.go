package workerpool

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

const size = 5

func double(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func TestNew(t *testing.T) {
	type args[J any, R any] struct {
		size   int
		worker func(id int, jobs <-chan int, results chan<- int)
	}
	tests := []struct {
		name string
		args args[int, func(id int, jobs <-chan int, results chan<- int)]
	}{
		{
			name: "test construction",
			args: args[int, func(id int, jobs <-chan int, results chan<- int)]{
				size:   5,
				worker: double,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := New(tt.args.size, tt.args.worker)
			if wp == nil {
				t.Error("New() = nil, want not nil")
			}
			if wp.Size != size {
				t.Errorf("New() Size = %d, want %d", wp.Size, size)
			}
		})
	}
}

func TestPool_Run(t *testing.T) {
	type args struct {
		jobs []int
	}
	tests := []struct {
		name        string
		args        args
		wantResults []int
	}{
		{
			name:        "test double integers",
			args:        args{jobs: []int{1, 2, 3, 4, 5}},
			wantResults: []int{2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := New(5, double)
			gotResults := wp.Run(tt.args.jobs)
			sort.Ints(gotResults)
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("Run() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}
