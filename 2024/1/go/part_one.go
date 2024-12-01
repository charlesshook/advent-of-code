package main

import (
	"math"
)

func partOne(firstList *[]int, secondList *[]int) int {
	var answer float64
	for i := 0; i < len(*firstList); i++ {
		answer += math.Abs(float64((*firstList)[i]) - float64((*secondList)[i]))
	}

	return int(answer)
}
