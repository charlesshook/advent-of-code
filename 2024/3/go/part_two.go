package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func partTwo(scanner *bufio.Scanner) int {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		pattern := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)|don't\\(\\)|do\\(\\)")
		validMatches := pattern.FindAllString(line, -1)

		for i := 0; i < len(validMatches); i++ {
			validMatches[i] = strings.Trim(validMatches[i], "mul()")
		}

		dontSeen := false

		for i := 0; i < len(validMatches); i++ {
			splitStrings := strings.Split(validMatches[i], ",")

			if splitStrings[0] == "don't" {
				dontSeen = true
			} else if splitStrings[0] == "do" {
				dontSeen = false
			} else if !dontSeen {
				firstValue, _ := strconv.Atoi(splitStrings[0])
				secondValue, _ := strconv.Atoi(splitStrings[1])

				total += (firstValue * secondValue)
			}
		}

	}

	return total
}
