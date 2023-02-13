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

func solveStep(grid [][]int) [][]int {
	return humanSolve(grid)
}

func autoSolve(grid [][]int) [][]int {
	edges := make(map[string]bool)
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			for _, adjCoord := range getValidAdjacentCells(grid, r, c) {
				if grid[adjCoord[0]][adjCoord[1]] != HIDDEN && grid[adjCoord[0]][adjCoord[1]] != FLAG {
					edges[fmt.Sprint("%d,%d", adjCoord[0], adjCoord[1])] = true
				}
			}
		}
	}
	// Game start, pick a random spot to begin
	if len(edges) == 0 {
		return [][]int{{int(rand.Float64() * float64(len(grid))), int(rand.Float64() * float64(len(grid[0])))}}
	}

	return nil
}

func humanSolve(grid [][]int) [][]int {
	fmt.Println("Input reveal coordinate seperated by commas (row,col):")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	coord := strings.Split(input, ",")
	r, _ := strconv.Atoi(coord[0])
	c, _ := strconv.Atoi(coord[1])
	reveal := []int{r, c}

	ret := make([][]int, 1)
	ret[0] = reveal
	fmt.Println("Input coordinates to flag seperated by commas (row,col), enter an empty entry to finish:")
	for scanner.Scan() {
		input = scanner.Text()
		if input == "" {
			break
		}

		coord = strings.Split(input, ",")
		r, _ = strconv.Atoi(coord[0])
		c, _ = strconv.Atoi(coord[1])
		ret = append(ret, []int{r, c})
	}

	return ret
}
