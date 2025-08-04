package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCustomWaitGroupAdd(t *testing.T) {
	nums := []int64{1, 5, 6, 9, 0, 1, -123, 8}

	for ti, num := range nums {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			wg := NewCustomWaitGroup()
			wg.Add(int(num))

			want := num
			if num < 0 {
				want = 0
			}

			if got := wg.waiter.Load(); got != want {
				t.Errorf("got: %d, want: %d", got, want)
			}
		})
	}
}

func TestCustomWaitGroupDone(t *testing.T) {
	nums := []int64{1, 5, 6, 9, 0, 1, -123, 8}

	for ti, num := range nums {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			wg := NewCustomWaitGroup()
			wg.waiter.Add(num) // проходим мимо wg.Add напрямую к значению

			if num < 0 {
				num = -num
			}
			for range num {
				wg.Done()
			}

			if got := wg.waiter.Load(); got != 0 {
				t.Errorf("got: %d, want: %d", got, 0)
			}
		})
	}
}

func TestCustomWaitGroupWait(t *testing.T) {
	t.Run("тест Wait()", func(t *testing.T) {
		wg := NewCustomWaitGroup()
		done := make(chan bool)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		wg.Add(1)

		go func() {
			wg.Wait()

			done <- true
		}()

		go func() {
			select {
			case <-done: // wg.Wait() не сработал
				return
			default:
				wg.Done()
			}
		}()

		select {
		case <-ctx.Done():
			t.Error("Wait() не закончил работу")
		case <-done: // ok!
		}
	})
}
