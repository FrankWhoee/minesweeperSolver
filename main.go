package main

import "fmt"

func main() {
	n := 9
	playerBoard := createGrid(n, n)
	solution := generateGrid(n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			playerBoard[r][c] = -1
		}
	}

	turn := 0
	for {
		printGrid(playerBoard)
		fmt.Printf("TURN: %d\n", turn)
		gameIsDone := true
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				if solution[r][c] != MINE && playerBoard[r][c] == HIDDEN {
					gameIsDone = false
				}
			}
		}
		if gameIsDone {
			fmt.Println("GAME OVER. YOU WIN.")
			fmt.Println("---------------------------------SOLUTION---------------------------------")
			printGrid(solution)
			return
		}
		actions := solveStep(playerBoard)
		if actions == nil {
			fmt.Println("Action can not be nil.")
			return
		}
		for i := 0; i < len(actions); i++ {
			if i == 0 {
				if playerBoard[actions[i][0]][actions[i][1]] == FLAG {
					fmt.Printf("Tried to reveal a flag.")
					return
				}

				// Generate a (sort of) playable board at the start
				if turn == 0 {
					for solution[actions[i][0]][actions[i][1]] != 0 {
						solution = generateGrid(n)
					}
				}
				if solution[actions[i][0]][actions[i][1]] == MINE {
					fmt.Printf("GAME OVER. HIT MINE AT %d,%d\n", actions[i][0], actions[i][1])
					playerBoard[actions[i][0]][actions[i][1]] = MINE
					fmt.Println("--------------------------------YOUR BOARD--------------------------------")
					printGrid(playerBoard)
					fmt.Println("---------------------------------SOLUTION---------------------------------")
					printGrid(solution)
					return
				}

				if playerBoard[actions[i][0]][actions[i][1]] == HIDDEN {
					recursiveFill(playerBoard, solution, actions[i][0], actions[i][1])
				} else {
					fmt.Println("Tried to reveal an already revealed tile.")
					return
				}
			} else {
				if playerBoard[actions[i][0]][actions[i][1]] == FLAG {
					fmt.Println("FLAG TO HIDDEN")
					playerBoard[actions[i][0]][actions[i][1]] = HIDDEN
				} else if playerBoard[actions[i][0]][actions[i][1]] == HIDDEN {
					fmt.Println("HIDDEN TO FLAG")
					playerBoard[actions[i][0]][actions[i][1]] = FLAG
				} else {
					fmt.Println("Tried to flag a revealed tile.")
					return
				}
			}
		}
		turn++
	}

}
