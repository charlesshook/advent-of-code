package main

import (
	"bufio"
	"strconv"
	"strings"
)

func partOne(scanner *bufio.Scanner) int {
	var safeCount int

	for scanner.Scan() {
		line := scanner.Text()
		results := strings.Split(line, " ")
		var lineSlice []int

		for i := 0; i < len(results); i++ {
			item, _ := strconv.Atoi(results[i])
			lineSlice = append(lineSlice, item)
		}

		var increased bool
		var decreased bool
		safe := true

		for i := 0; i < len(lineSlice)-1; i++ {
			difference := lineSlice[i+1] - lineSlice[i]

			if difference > 0 {
				increased = true
				if difference > 3 {
					safe = false
				}
			} else if difference < 0 {
				decreased = true
				if difference < -3 {
					safe = false
				}
			} else if difference == 0 {
				safe = false
			}
		}

		if increased && decreased {
			safe = false
		}

		if safe {
			safeCount += 1
		}
	}

	return int(safeCount)
}
