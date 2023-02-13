package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Takes in a grid of hidden tiles, and can choose to flag any number of tiles as mines, but can only choose one tile to reveal
// Returns a list of tiles. First one is to reveal, the rest is to flag.

func solveStep(grid [][]int) [2][][]int {
	return autoSolve(grid)
}

func autoSolve(grid [][]int) [2][][]int {
	edges := make(map[string]bool)
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			for _, adjCoord := range getValidAdjacentCells(grid, r, c) {
				if grid[adjCoord[0]][adjCoord[1]] != HIDDEN && grid[adjCoord[0]][adjCoord[1]] != FLAG {
					edges[coordToString(adjCoord)] = true
				}
			}
		}
	}
	// Game start, pick a random spot to begin
	if len(edges) == 0 {
		return [2][][]int{{{int(rand.Float64() * float64(len(grid))), int(rand.Float64() * float64(len(grid[0])))}}, {}}
	}

	flagSet := make(map[string]bool)
	for coord := range edges {
		split := strings.Split(coord, ",")
		r, _ := strconv.Atoi(split[0])
		c, _ := strconv.Atoi(split[1])
		var adjHiddens [][]int
		var adjFlags [][]int
		for _, adjCoord := range getValidAdjacentCells(grid, r, c) {
			if grid[adjCoord[0]][adjCoord[1]] == FLAG {
				adjFlags = append(adjFlags, adjCoord)
			} else if grid[adjCoord[0]][adjCoord[1]] == HIDDEN {
				adjHiddens = append(adjHiddens, adjCoord)
			}
		}
		if len(adjFlags)+len(adjHiddens) == grid[r][c] {
			for _, hidden := range adjHiddens {
				flagSet[coordToString(hidden)] = true
			}
		}
	}

	revealSet := make(map[string]bool)
	for coord := range edges {
		split := strings.Split(coord, ",")
		r, _ := strconv.Atoi(split[0])
		c, _ := strconv.Atoi(split[1])
		var adjHiddens [][]int
		var adjFlags [][]int
		for _, adjCoord := range getValidAdjacentCells(grid, r, c) {
			if grid[adjCoord[0]][adjCoord[1]] == FLAG || flagSet[fmt.Sprintf("%d,%d", adjCoord[0], adjCoord[1])] {
				adjFlags = append(adjFlags, adjCoord)
			} else if grid[adjCoord[0]][adjCoord[1]] == HIDDEN {
				adjHiddens = append(adjHiddens, adjCoord)
			}
		}
		if len(adjFlags) == grid[r][c] {
			for _, hidden := range adjHiddens {
				revealSet[coordToString(hidden)] = true
			}
		}
	}

	flags := mapToCoords(flagSet)
	reveals := mapToCoords(revealSet)
	fmt.Println(reveals)
	fmt.Println(flags)

	//if len(revealSet) == 0 {
	//	return autoSolve(grid)
	//}

	return [2][][]int{reveals, flags}
}

func coordToString(hidden []int) string {
	return fmt.Sprintf("%d,%d", hidden[0], hidden[1])
}

func mapToCoords(inputMap map[string]bool) [][]int {
	var coords [][]int
	for coord := range inputMap {
		split := strings.Split(coord, ",")
		r, _ := strconv.Atoi(split[0])
		c, _ := strconv.Atoi(split[1])
		coords = append(coords, []int{r, c})
	}
	return coords
}

func humanSolve(grid [][]int) [2][][]int {
	fmt.Println("Input coordinates to flag seperated by commas (row,col), enter an empty entry to finish:")
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	var flags [][]int
	for scanner.Scan() {
		input = scanner.Text()
		if input == "" {
			break
		}

		coord := strings.Split(input, ",")
		r, _ := strconv.Atoi(coord[0])
		c, _ := strconv.Atoi(coord[1])
		flags = append(flags, []int{r, c})
	}

	fmt.Println("Input coordinates to reveal seperated by commas (row,col), enter an empty entry to finish:")
	var reveals [][]int
	for scanner.Scan() {
		input = scanner.Text()
		if input == "" {
			break
		}

		coord := strings.Split(input, ",")
		r, _ := strconv.Atoi(coord[0])
		c, _ := strconv.Atoi(coord[1])
		reveals = append(reveals, []int{r, c})
	}
	return [2][][]int{reveals, flags}
}
