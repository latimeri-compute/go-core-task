package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sync"
)

func main() {
	numbers := []uint8{5, 34, 27, 90, 12}
	fmt.Printf("исходный слайс чисел: %v\n", numbers)
	chuint := make(chan uint8)
	chfloat := make(chan float64)

	done := make(chan struct{})
	var wg sync.WaitGroup

	go pipe(chuint, chfloat)

	go func() {
		printFromChan(os.Stdout, chfloat)
		close(done)
	}()

	for _, n := range numbers {
		wg.Add(1)
		go func(n uint8) {
			defer wg.Done()
			chuint <- n
		}(n)
	}

	wg.Wait()
	close(chuint)
	<-done
}

func printFromChan[T any](w io.Writer, ch chan T) {
	for n := range ch {
		fmt.Fprintf(w, "%v\n", n)
	}
}

func pipe(sender chan uint8, receiver chan float64) {
	defer close(receiver)
	for num := range sender {
		receiver <- math.Pow(float64(num), 3)
	}
}
