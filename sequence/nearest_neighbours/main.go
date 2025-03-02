package main

import (
	"fmt"
	"math"
	"slices"
)

func nearestNeighboursScore(seats []int, people []int) int {
	slices.Sort(seats)

	nns, numSeats := 0, len(seats)

	for _, ppl := range people {
		iseat, isFound := slices.BinarySearch(seats, ppl)
		if isFound {
			continue
		}

		if iseat == numSeats { // index out of range, substitute with maximal seat
			iseat -= 1
		}

		diff := math.Abs((float64)(ppl - seats[iseat]))

		if iseat != 0 { // choose between selected and earlier seats
			diff = math.Min(diff, math.Abs((float64)(ppl-seats[iseat-1])))
		}

		nns += (int)(diff)
	}

	return nns
}

func main() {
	fmt.Println(nearestNeighboursScore([]int{9, 1, 8, 2, 7, 4}, []int{0, 2, 5, 11}))
}
