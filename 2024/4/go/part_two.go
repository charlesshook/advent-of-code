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

	for i := range len(moves) {
		for j := i + 1; j < len(moves); j++ {
			if pointsIntersect(moves[i][0], moves[i][1], moves[j][0], moves[j][1]) {
				fmt.Println(moves[i][0], moves[i][1], moves[j][0], moves[j][1])
				count += 1
			}
		}
	}

	return count
}
