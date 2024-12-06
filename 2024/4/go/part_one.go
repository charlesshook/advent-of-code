package main

import (
	"bufio"
)

func partOne(scanner *bufio.Scanner) int {
	var charArray [][]rune
	words := []string{"XMAS"}

	for scanner.Scan() {
		line := scanner.Text()
		charArray = append(charArray, []rune(line))
	}

	return len(solve(charArray, words))
}
