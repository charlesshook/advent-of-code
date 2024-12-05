package main

import (
	"bufio"
)

func partOne(scanner *bufio.Scanner) int {
	var charArray [][]rune
	words := []string{"XMAS"}

	for scanner.Scan() {
		line := scanner.Text()
		charArray = append(charArray, []rune(line))
	}

	return len(solve(charArray, words))
}

type Direction struct {
	x, y int
}

type Node struct {
	children map[rune]*Node
	isWord   bool
}

type Trie struct {
	root *Node
}

func newTrie() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}

func (t *Trie) insert(word string) {
	node := t.root
	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			node.children[r] = &Node{children: make(map[rune]*Node)}
		}
		node = node.children[r]
	}
	node.isWord = true
}

func (t *Trie) search(word string) bool {
	node := t.root

	for _, r := range word {
		if _, ok := node.children[r]; !ok {
			return false
		}

		node = node.children[r]
	}

	return node.isWord
}

func solve(grid [][]rune, words []string) []Direction {
	n, m := len(grid), len(grid[0])
	trie := newTrie()
	moves := []Direction{}
	for _, word := range words {
		trie.insert(word)
	}

	directions := []Direction{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	for i := range n {
		for j := range m {
			if _, ok := trie.root.children[grid[i][j]]; ok {
				for _, direction := range directions {
					check(grid, trie.root, i, j, direction.x, direction.y, &moves)
				}
			}
		}
	}

	return moves
}

func check(grid [][]rune, trie *Node, i int, j int, xDirection int, yDirection int, moves *[]Direction) {
	n, m := len(grid), len(grid[0])
	node := trie

	for i >= 0 && i < n && j >= 0 && j < m {
		char := grid[i][j]

		if nextNode, ok := node.children[char]; ok {
			node = nextNode

			if node.isWord {
				*moves = append(*moves, Direction{i, j})
			}
		} else {
			break
		}

		i += xDirection
		j += yDirection
	}
}
