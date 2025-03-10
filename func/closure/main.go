package main

import (
	"fmt"
)

func main() {
	a, b := 0, 1

	fibo := func() int { // anonymous function
		a, b = b, a+b
		return b
	}

	fmt.Println("fubo: ", fibo(), fibo(), fibo(), fibo(), fibo())

	fibo2 := func(a, b int) int { // anonymous function
		a, b = b, a+b
		return b
	}

	fmt.Println("fibo2:", fibo2(0, 1), fibo2(1, 1), fibo2(1, 2), fibo2(2, 3), fibo2(3, 5))

	fibo3, fibo4 := mkFiboSeq(0, 1), mkFiboSeq(3, 5)
	fmt.Println("fibo3:", fibo3(), fibo3(), fibo3(), fibo3(), fibo3())
	fmt.Println("fibo4:", fibo4(), fibo4(), fibo4(), fibo4(), fibo4())
}

func mkFiboSeq(a, b int) func() int {
	return func() int { // closure
		a, b = b, a+b
		return b
	}
}
