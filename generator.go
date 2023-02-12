package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printGrid(grid [][]int) {
	for c := 0; c < len(grid); c++ {
		fmt.Println(grid[c])
	}
}

const MINE = 9

// naive generator, does not check for uniqueness of solution
func generateGrid(n int) [][]int {
	rand.Seed(time.Now().UnixNano())
	cols := n
	rows := n
	grid := make([][]int, cols)
	for i := 0; i < cols; i++ {
		grid[i] = make([]int, rows)
	}

	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			if rand.Float32() < 0.2 {
				grid[c][r] = MINE
			}
		}
	}

	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			if grid[c][r] == 0 {
				count := 0
				for dc := -1; dc <= 1; dc++ {
					for dr := -1; dr <= 1; dr++ {
						if (dc != 0 || dr != 0) && c+dc >= 0 && c+dc < cols && r+dr >= 0 && r+dr < rows && grid[c+dc][r+dr] == MINE {
							count += 1
						}
					}
				}
				grid[c][r] = count
			}
		}
	}

	return grid
}
