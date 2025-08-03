package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestPipe(t *testing.T) {
	tests := []struct {
		num  uint8
		want float64
	}{
		{
			num:  3,
			want: 27,
		},
		{
			num:  1,
			want: 1,
		},
		{
			num:  2,
			want: 8,
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			chuint := make(chan uint8)
			chfloat := make(chan float64)
			var wg sync.WaitGroup

			go pipe(chuint, chfloat)

			go func(num uint8) {
				chuint <- num
			}(test.num)

			wg.Add(1)
			var got float64
			go func() {
				defer wg.Done()
				got = <-chfloat
				close(chfloat)
			}()

			wg.Wait()
			if test.want != got {
				t.Errorf("got: %v; want: %v", got, test.want)
			}
		})
	}

}

func TestPrintFromChan(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{
			nums: []int{87, 43, 1, 90, 123, 63, 123},
		},
		{
			nums: []int{123123, 94535, 004, 1236, 2342, 3, 8},
		},
	}
	for ti, test := range tests {
		t.Run(fmt.Sprintf("%02d", ti), func(t *testing.T) {
			var wg sync.WaitGroup
			ch := make(chan int)
			buf := bytes.NewBuffer([]byte{})

			go func() {
				defer wg.Done()
				printFromChan(buf, ch)
			}()

			wg.Add(len(test.nums))
			for _, num := range test.nums {
				go func(num int) {
					defer wg.Done()
					ch <- num
				}(num)
			}
			wg.Wait()
			wg.Add(1) // нужно, чтобы printFromChan закончил работу
			close(ch)
			wg.Wait()

			got := stringToIntArr(t, buf.String())

			slices.Sort(test.nums)
			slices.Sort(got)

			if !slices.Equal(test.nums, got) {
				t.Errorf("got: %v; want: %v", got, test.nums)
			}
		})
	}
}

func stringToIntArr(t *testing.T, s string) []int {
	var arr []int
	res := strings.SplitSeq(s, "\n")
	for item := range res {
		if len(item) < 1 {
			continue
		}
		n, err := strconv.Atoi(item)
		if err != nil {
			t.Fatal(err)
		}
		arr = append(arr, n)
	}

	return arr
}
