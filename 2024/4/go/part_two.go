package main

import (
	"bufio"
)

func partTwo(scanner *bufio.Scanner) int {
	var charArray [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		charArray = append(charArray, []rune(line))
	}

	count := 0

	height := len(charArray)
	width := len(charArray[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i+2 < height && j+2 < width && charArray[i][j] == 'M' && charArray[i][j+2] == 'M' && charArray[i+1][j+1] == 'A' && charArray[i+2][j] == 'S' && charArray[i+2][j+2] == 'S' {
				count += 1
			} else if i+2 < height && j+2 < width && charArray[i][j] == 'S' && charArray[i][j+2] == 'M' && charArray[i+1][j+1] == 'A' && charArray[i+2][j] == 'S' && charArray[i+2][j+2] == 'M' {
				count += 1
			} else if i+2 < height && j+2 < width && charArray[i][j] == 'S' && charArray[i][j+2] == 'S' && charArray[i+1][j+1] == 'A' && charArray[i+2][j] == 'M' && charArray[i+2][j+2] == 'M' {
				count += 1
			} else if i+2 < height && j+2 < width && charArray[i][j] == 'M' && charArray[i][j+2] == 'S' && charArray[i+1][j+1] == 'A' && charArray[i+2][j] == 'M' && charArray[i+2][j+2] == 'S' {
				count += 1
			}
		}
	}

	return count
}
