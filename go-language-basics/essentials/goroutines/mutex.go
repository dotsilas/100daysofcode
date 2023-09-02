package goroutines

import (
	"fmt"
	"sync"
)

type sum struct {
	mu sync.Mutex
	sum int
}

func (s *sum) Get() int {
	// locks method at the start
	s.mu.Lock()
	// unlock at the end, (defer executes at the end)
	defer s.mu.Unlock()
	return s.sum
}

func (s *sum) Add(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sum++
}

func mutex() {
	mySum := &sum{}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++{
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			mySum.Add(x)
		}(i)
	}

	result := mySum.Get()
	fmt.Printf("%d", result)
}
