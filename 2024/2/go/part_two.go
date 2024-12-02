package main

import (
	"bufio"
	"fmt"
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

		var increased bool
		var decreased bool
		safe := true
		unSafeCount := 0

		for i := 0; i < len(lineSlice)-1; i++ {
			difference := lineSlice[i+1] - lineSlice[i]

			if difference > 0 {
				increased = true
				if difference > 3 {
					safe = false
					unSafeCount += 1
				}
			} else if difference < 0 {
				decreased = true
				if difference < -3 {
					safe = false
					unSafeCount += 1
				}
			} else if difference == 0 {
				safe = false
				unSafeCount += 1
			}
		}

		if increased && decreased {
			safe = false
		}

		if safe {
			safeCount += 1
		} else if !safe {
			if safeCount <= 1 {
				safeCount += 1
				safe = true
			}
		}

		fmt.Println(safe)
	}

	return int(safeCount)
}
