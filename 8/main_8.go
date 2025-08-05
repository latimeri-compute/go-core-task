package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CustomWaitGroup struct {
	waiter atomic.Int64
	cond   *sync.Cond
	mu     sync.Mutex
}

func NewCustomWaitGroup() *CustomWaitGroup {
	wg := &CustomWaitGroup{}
	wg.cond = sync.NewCond(&wg.mu)
	return wg
}

func (wg *CustomWaitGroup) Add(delta int) {
	if wg.waiter.Add(int64(delta)) <= 0 {
		wg.mu.Lock()
		defer wg.mu.Unlock()
		wg.cond.Broadcast()
		// это, наверное, одно из самых явных отличий от стандартной библиотеки
		if wg.waiter.Load() < 0 {
			wg.waiter.Store(0)
		}
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	for wg.waiter.Load() > 0 {
		wg.cond.Wait()
	}
}

func main() {
	nums := make([]int, 0)

	for i := range 50 {
		nums = append(nums, i)
	}

	wg := NewCustomWaitGroup()
	var l atomic.Int32

	for n := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
			l.Add(1)
		}(n)
	}

	wg.Wait()

	fmt.Println("Printing done!")
	fmt.Printf("Total lines: %d\n", l.Load())
}
