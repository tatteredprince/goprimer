package main

import (
	"fmt"
)

func fibo(n int) <-chan int {
	out := make(chan int)
	if n < 3 {
		panic("wrong sequence length")
	}
	go func(n int) {
		a, b := 0, 1
		out <- a
		out <- b
		for range n - 2 {
			a, b = b, a+b
			out <- b
		}
		close(out)
	}(n)
	return out
}

func main() {
	ch := fibo(5)
	for range 5 {
		fmt.Println(<-ch)
	}

	for n := range fibo(10) {
		fmt.Println(n)
	}
}
