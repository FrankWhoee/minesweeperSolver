package main

import "fmt"

func printGrid(grid [][]int) {
	for r := 0; r < len(grid); r++ {
		fmt.Println(grid[r])
	}
}

func createGrid(cols int, rows int) [][]int {
	grid := make([][]int, cols)
	for r := 0; r < len(grid); r++ {
		grid[r] = make([]int, rows)
	}
	return grid
}

func recursiveFill(hidden [][]int, solution [][]int, startRow int, startCol int) {
	hidden[startRow][startCol] = solution[startRow][startCol]
	if solution[startRow][startCol] == 0 {
		rows := len(solution)
		cols := len(solution[0])
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if (dc != 0 || dr != 0) && startCol+dc >= 0 && startCol+dc < cols && startRow+dr >= 0 && startRow+dr < rows {
					if hidden[startRow+dr][startCol+dc] == -1 {
						recursiveFill(hidden, solution, startRow+dr, startCol+dc)
					}
				}
			}
		}
	}
}
