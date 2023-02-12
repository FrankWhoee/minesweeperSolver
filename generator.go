package main

import (
	"math/rand"
)

// naive generator, does not check for uniqueness of solution
func generateGrid(n int) [][]int {
	//rand.Seed(time.Now().UnixNano())
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

				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if (dc != 0 || dr != 0) && c+dc >= 0 && c+dc < cols && r+dr >= 0 && r+dr < rows && grid[r+dr][c+dc] == MINE {
							count += 1
						}
					}
				}
				grid[r][c] = count
			}
		}
	}

	return grid
}
