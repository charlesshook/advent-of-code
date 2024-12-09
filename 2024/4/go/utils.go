package main

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
	pathArray := []Point{}

	for i >= 0 && i < n && j >= 0 && j < m {
		char := grid[i][j]

		if nextNode, ok := node.children[char]; ok {
			node = nextNode
			pathArray = append(pathArray, Point{i, j})

			if node.isWord {
				*moves = append(*moves, pathArray)
				pathArray = pathArray[:0]
			}
		} else {
			pathArray = pathArray[:0]
			break
		}

		i += xPoint
		j += yPoint
	}
}
