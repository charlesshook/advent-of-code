package main

import (
	"bufio"
	"strconv"
	"strings"
)

func partTwo(scanner *bufio.Scanner) int {
	var safeCount int

	for scanner.Scan() {
		line := scanner.Text()
		results := strings.Split(line, " ")
		var lineSlice []int

		for i := 0; i < len(results); i++ {
			item, _ := strconv.Atoi(results[i])
			lineSlice = append(lineSlice, item)
		}

		if isSafe(lineSlice) {
			safeCount += 1
		} else {
			for i := 0; i < len(lineSlice); i++ {
				if isSafe(append(append([]int{}, lineSlice[:i]...), lineSlice[i+1:]...)) {
					safeCount += 1
					break
				}
			}
		}

	}

	return safeCount
}

func isSafe(lineSlice []int) bool {
	var increased bool
	var decreased bool

	for i := 0; i < len(lineSlice)-1; i++ {
		difference := lineSlice[i+1] - lineSlice[i]

		if difference > 3 || difference < -3 || difference == 0 {
			return false
		}

		if difference > 0 {
			increased = true
		} else if difference < 0 {
			decreased = true
		}
	}

	if increased && decreased {
		return false
	}

	return !(increased && decreased)
}
