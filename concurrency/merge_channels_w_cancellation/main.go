package main

import (
	"fmt"
	"sync"
)

func merge(done <-chan struct{}, channels ...<-chan int) <-chan int {
	out := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(len(channels))

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for value := range ch {
				select {
				case out <- value:
				case <-done:
					return
				}
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
	done := make(chan struct{})

	fanOut := func(value, count int) <-chan int {
		out := make(chan int)
		go func(value, count int) {
			defer close(out)
			for range count {
				select {
				case <-done:
					return
				default:
					out <- value
				}
			}
		}(value, count)
		return out
	}

	fanIn := merge(done, fanOut(1, 10), fanOut(2, 10), fanOut(3, 10))

	readAndCutOff := func(nRead, nCutoff int) {
		for range nRead {
			fmt.Println(<-fanIn)
		}
		for range nCutoff {
			done <- struct{}{}
		}
	}

	readAndCutOff(4, 1)
	readAndCutOff(3, 1)
	readAndCutOff(5, 0)
	close(done)

	<-fanIn // clear channel from residue data

	val, ok := <-fanIn // test if channel is closed
	if !ok {
		fmt.Println("Fan-in channel closed")
	} else { // never gonna reach here
		fmt.Println("Fan-in channel got value:", val)
	}
}
