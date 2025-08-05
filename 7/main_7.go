package main

import "sync"

func mergeChannels[T any](channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	merged := make(chan T)

	out := func(c <-chan T) {
		for item := range c {
			merged <- item
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go out(c)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged

}
