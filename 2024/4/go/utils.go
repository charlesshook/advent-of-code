package main

import "fmt"

func print2DArray(charArray [][]rune) {
	for i := 0; i < len(charArray); i++ {
		for j := 0; j < len(charArray); j++ {
			fmt.Printf("%c", charArray[i][j])
		}
		fmt.Println()
	}
}
