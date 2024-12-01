package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parseSlices(scanner *bufio.Scanner) ([]int, []int) {
	var firstList []int
	var secondList []int

	for scanner.Scan() {
		line := scanner.Text()
		results := strings.Split(line, "   ")

		itemone, _ := strconv.Atoi(results[0])
		itemtwo, err := strconv.Atoi(results[1])

		if err != nil {
			fmt.Println(err)
		}

		firstList = append(firstList, itemone)
		secondList = append(secondList, itemtwo)
	}

	return firstList, secondList
}
