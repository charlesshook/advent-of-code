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

type Point struct {
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

func solve(grid [][]rune, words []string) [][]Point {
	n, m := len(grid), len(grid[0])
	trie := newTrie()
	moves := [][]Point{}
	for _, word := range words {
		trie.insert(word)
	}

	directions := []Point{
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

func check(grid [][]rune, trie *Node, i int, j int, xPoint int, yPoint int, moves *[][]Point) {
	n, m := len(grid), len(grid[0])
	node := trie

	// startingI, startingJ := i, j
	pathArray := []Point{}
	why := ""

	for i >= 0 && i < n && j >= 0 && j < m {
		char := grid[i][j]

		if nextNode, ok := node.children[char]; ok {
			node = nextNode
			why += fmt.Sprintf("(%d, %d)", i, j)
			pathArray = append(pathArray, Point{i, j})

			if node.isWord {
				*moves = append(*moves, pathArray)
				pathArray = pathArray[:0]
				why = ""
			}
		} else {
			pathArray = pathArray[:0]
			break
		}

		i += xPoint
		j += yPoint
	}
}

func onSegment(p Point, q Point, r Point) bool {
	if q.x <= max(p.x, r.x) && q.x >= min(p.x, r.x) && q.y <= max(p.y, r.y) && p.y >= min(p.y, r.y) {
		return true
	}

	return false
}

func orientation(p Point, q Point, r Point) int {
	val := (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)

	if val == 0 {
		return 0
	} else if val > 0 {
		return 1
	} else {
		return 2
	}
}

func pointsIntersect(p1 Point, q1 Point, p2 Point, q2 Point) bool {
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	if o1 != o2 && o3 != o4 {
		return true
	} else if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	} else if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	} else if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	} else if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	} else {
		return false
	}
}
