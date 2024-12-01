package main

import "sort"

func partTwo(firstList *[]int, secondList *[]int) int {
	sort.Ints(*firstList)
	sort.Ints(*secondList)
	var answer int
	secondListMap := make(map[int]int)

	// Setup secondListMap
	for i := 0; i < len(*secondList); i++ {
		secondListMap[(*secondList)[i]] = secondListMap[(*secondList)[i]] + 1
	}

	for i := 0; i < len(*firstList); i++ {
		answer += (*firstList)[i] * secondListMap[(*firstList)[i]]
	}

	return answer
}
