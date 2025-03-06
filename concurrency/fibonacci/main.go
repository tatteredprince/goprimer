package main

import (
	"fmt"
)

func fibo(a, b, n int) <-chan int {
	out := make(chan int)
	go func(a, b, n int) {
		out <- a
		out <- b
		for i := 0; i < n; i++ {
			a, b = b, a+b
			out <- b
		}
		close(out)
	}(a, b, n)
	return out
}

func main() {
	ch := fibo(0, 1, 5)
	fmt.Print(<-ch)

	for n := range ch {
		fmt.Print(" ", n)
	}
	fmt.Print("\n")

	ch = fibo(10, 11, 3)
	fmt.Println(<-ch, <-ch, <-ch, <-ch, <-ch)
}
