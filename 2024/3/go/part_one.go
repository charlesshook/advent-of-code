package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func partOne(scanner *bufio.Scanner) int {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		pattern := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
		validMatches := pattern.FindAllString(line, -1)

		for i := 0; i < len(validMatches); i++ {
			validMatches[i] = strings.Trim(validMatches[i], "mul()")
		}

		for i := 0; i < len(validMatches); i++ {
			splitStrings := strings.Split(validMatches[i], ",")

			firstValue, _ := strconv.Atoi(splitStrings[0])
			secondValue, _ := strconv.Atoi(splitStrings[1])

			total += (firstValue * secondValue)
		}

	}

	return total
}
