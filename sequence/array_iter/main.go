package main

import "fmt"

func main() {
	a, b := [...]int{1, 2, 3}, [3]int{0: 4, 5, 2: 6}
	fmt.Printf("types a = %T, b = %T\n", a, b)

	for i, v := range a { // iterate over implicit copy of the original array
		if i == 1 {
			a = b // assign elements to the original array
		}
		fmt.Println(v)
	}

	fmt.Printf("a = %v, b = %v\n", a, b)

	a, b = [3]int{1, 2, 3}, [3]int{4, 5, 6}

	for i, v := range &a { // iterate over the original array
		if i == 1 {
			a = b // assign elements to the original array
		}
		fmt.Println(v)
	}

	fmt.Printf("a = %v, b = %v\n", a, b)
}
