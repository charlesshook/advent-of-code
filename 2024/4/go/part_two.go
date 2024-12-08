package main

import (
	"bufio"
	"fmt"
)

func partTwo(scanner *bufio.Scanner) int {
	var charArray [][]rune
	words := []string{"MAS"}

	for scanner.Scan() {
		line := scanner.Text()
		charArray = append(charArray, []rune(line))
	}

	moves := solve(charArray, words)
	count := 0

	fmt.Println(moves)

	for i := 0; i < len(moves); i++ {
		for j := 1; j < len(moves); j++ {
			if pointsIntersect(moves[i][0], moves[i][2], moves[j][0], moves[j][2]) {
				count += 1
			}
		}
	}

	return count
}
