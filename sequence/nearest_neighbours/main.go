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
		iSeat, isFound := slices.BinarySearch(seats, ppl)
		if isFound {
			continue
		}

		if iSeat == numSeats { // index out of range, substitute with maximal seat
			iSeat -= 1
		}

		diff := math.Abs((float64)(ppl - seats[iSeat]))

		if iSeat != 0 { // choose between selected and earlier seats
			diff = math.Min(diff, math.Abs((float64)(ppl-seats[iSeat-1])))
		}

		nns += (int)(diff)
	}

	return nns
}

func main() {
	fmt.Println(nearestNeighboursScore([]int{9, 1, 8, 2, 7, 4}, []int{0, 2, 5, 11}))
}
