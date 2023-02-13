package main

import "fmt"

func printGrid(grid [][]int) {
	fmt.Printf("   ")
	for c := 0; c < len(grid[0]); c++ {
		fmt.Printf("%d  ", c)
	}
	fmt.Printf("\n")
	for r := 0; r < len(grid); r++ {
		fmt.Printf("%d  ", r)
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == FLAG {
				fmt.Printf("◈")
			} else if grid[r][c] == MINE {
				fmt.Printf("✸")
			} else if grid[r][c] == HIDDEN {
				fmt.Printf("☐")
			} else if grid[r][c] == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d", grid[r][c])
			}
			fmt.Printf("  ")
		}
		fmt.Printf("‎ \n")
	}
}

func createGrid(cols int, rows int) [][]int {
	grid := make([][]int, cols)
	for r := 0; r < len(grid); r++ {
		grid[r] = make([]int, rows)
	}
	return grid
}

func getValidAdjacentCells(grid [][]int, row int, col int) [][]int {
	ret := make([][]int, 0)
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if (dc != 0 || dr != 0) && col+dc >= 0 && col+dc < len(grid[0]) && row+dr >= 0 && row+dr < len(grid) {
				ret = append(ret, []int{row + dr, col + dc})
			}
		}
	}
	return ret
}

func recursiveFill(hidden [][]int, solution [][]int, startRow int, startCol int) {
	hidden[startRow][startCol] = solution[startRow][startCol]
	if solution[startRow][startCol] == 0 {
		for _, adjCoord := range getValidAdjacentCells(solution, startRow, startCol) {
			if hidden[adjCoord[0]][adjCoord[1]] == HIDDEN {
				recursiveFill(hidden, solution, adjCoord[0], adjCoord[1])
			}

		}
	}
}
