package main

import (
	"fmt"
	"sync"
)

func merge(channels ...<-chan int) <-chan int { // aka 'Fan-in'
	out := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(len(channels))

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for value := range ch {
				out <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	fanOut := func(value, count int) <-chan int {
		out := make(chan int)
		go func() {
			for range count {
				out <- value
			}
			close(out)
		}()
		return out
	}

	for n := range merge(fanOut(5, 3), fanOut(1, 2), fanOut(7, 5)) {
		fmt.Println(n)
	}
}
