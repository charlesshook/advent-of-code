package main

func partTwo(firstList *[]int, secondList *[]int) int {
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
