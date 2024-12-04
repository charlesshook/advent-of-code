package main

import (
	"bufio"
	"regexp"
	"strconv"
)

func partTwo(scanner *bufio.Scanner) int {
	total := 0
	allText := ""

	for scanner.Scan() {
		allText += scanner.Text()
	}

	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)
	validMatches := pattern.FindAllStringSubmatch(allText, -1)

	dontSeen := false

	for _, validMatch := range validMatches {
		if len(validMatch) > 0 {
			if validMatch[0] == "don't()" {
				dontSeen = true
			} else if validMatch[0] == "do()" {
				dontSeen = false
			} else if !dontSeen {
				firstValue, _ := strconv.Atoi(validMatch[1])
				secondValue, _ := strconv.Atoi(validMatch[2])

				total += (firstValue * secondValue)
			}
		}
	}

	return total
}
