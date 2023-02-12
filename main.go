package main

import "fmt"

func main() {
	n := 9
	hidden := createGrid(n, n)
	solution := generateGrid(n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			hidden[r][c] = -1
		}
	}

	actions := solveStep(hidden)
	for i := 0; i < len(actions); i++ {
		if i == 1 {
			if solution[actions[i][0]][actions[i][1]] == MINE {
				fmt.Printf("GAME OVER. HIT MINE AT %d,%d\n", actions[i][0], actions[i][1])
				hidden[actions[i][0]][actions[i][1]] = MINE
				fmt.Println("--------------------------------YOUR BOARD--------------------------------")
				printGrid(hidden)
				fmt.Println("---------------------------------SOLUTION---------------------------------")
				printGrid(solution)
				return
			} else if hidden[actions[i][0]][actions[i][1]] == HIDDEN {
				recursiveFill(hidden, solution, actions[i][0], actions[i][1])
			} else {
				fmt.Println("Tried to reveal an already revealed tile.")
				return
			}
		} else {
			if hidden[actions[i][0]][actions[i][1]] == FLAG {
				hidden[actions[i][0]][actions[i][1]] = HIDDEN
			} else if hidden[actions[i][0]][actions[i][1]] == HIDDEN {
				hidden[actions[i][0]][actions[i][1]] = FLAG
			} else {
				fmt.Println("Tried to flag a revealed tile.")
				return
			}
		}
	}
}
