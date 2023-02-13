package main

import (
	"fmt"
	"math/rand"
)

// naive generator, does not check for uniqueness of solution
func generateGrid(n int) [][]int {
	//seed := time.Now().UnixNano()
	seed := 1676267509194297249
	fmt.Println(seed)
	rand.Seed(int64(seed))
	rows := n
	cols := n
	grid := createGrid(rows, cols)

	for r := 0; r < rows; r++ {
		for c := 0; c < rows; c++ {
			if rand.Float32() < 0.2 {
				grid[r][c] = MINE
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < rows; c++ {
			if grid[r][c] == 0 {
				count := 0
				for _, adjCoord := range getValidAdjacentCells(grid, r, c) {
					if grid[adjCoord[0]][adjCoord[1]] == MINE {
						count += 1
					}
				}
				grid[r][c] = count
			}
		}
	}

	return grid
}
