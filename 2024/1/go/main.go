package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Run challenge part 1 or part 2.")
	flag.Parse()
	fmt.Printf("Running challenge part %d.\n", part)

	input, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("Could not open input.txt file. Error: %s.\n", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	firstList, secondList := parseSlices(scanner)

	if err = scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %s.\n", err)
	}

	if part == 1 {
		answer := partOne(&firstList, &secondList)
		fmt.Printf("The solution to part 1 is: %d\n", answer)
	} else if part == 2 {
		answer := partTwo(&firstList, &secondList)
		fmt.Printf("The solution to part 2 is: %d\n", answer)
	} else {
		fmt.Printf("Part %d is not a valid option.\n", part)
	}
}
