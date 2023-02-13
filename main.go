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
		if actions[0] == nil {
			fmt.Println("Reveals can not be nil.")
			return
		}

		for i, revealCoord := range actions[0] {
			// Generate a (sort of) playable board at the start
			if turn == 0 && i == 0 {
				for solution[revealCoord[0]][revealCoord[1]] != 0 {
					solution = generateGrid(n)
				}
			}

			if playerBoard[revealCoord[0]][revealCoord[1]] == FLAG {
				fmt.Printf("Tried to reveal a flag.")
				return
			}

			if solution[revealCoord[0]][revealCoord[1]] == MINE {
				fmt.Printf("GAME OVER. HIT MINE AT %d,%d\n", revealCoord[0], revealCoord[1])
				playerBoard[revealCoord[0]][revealCoord[1]] = MINE
				fmt.Println("--------------------------------YOUR BOARD--------------------------------")
				printGrid(playerBoard)
				fmt.Println("---------------------------------SOLUTION---------------------------------")
				printGrid(solution)
				return
			}

			if playerBoard[revealCoord[0]][revealCoord[1]] == HIDDEN {
				recursiveFill(playerBoard, solution, revealCoord[0], revealCoord[1])
			} else {
				fmt.Printf("Tried to reveal an already revealed tile: (%d,%d) \n", revealCoord[0], revealCoord[1])
				//return
			}
		}
		for _, flagCoord := range actions[1] {
			if playerBoard[flagCoord[0]][flagCoord[1]] == FLAG {
				playerBoard[flagCoord[0]][flagCoord[1]] = HIDDEN
			} else if playerBoard[flagCoord[0]][flagCoord[1]] == HIDDEN {
				playerBoard[flagCoord[0]][flagCoord[1]] = FLAG
			} else {
				fmt.Println("Tried to flag a revealed tile.")
				return
			}
		}
		turn++
	}
}
