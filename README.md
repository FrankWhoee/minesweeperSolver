# minesweeperSolver
I was playing minesweeper, and I was at the same time getting bored of playing minesweeper.

I automated it.

You can use the game framework to make your own minesweeper solver, just write your own `solver.go` file. The game driver `main.go` is expecting a function called `solveStep` that accepts a grid (`[][]int`) and outputs a type (`[2][][]int`) with the first list being a list of cells that you want to reveal, and the second list being a list of cells that you want to flag. **The solver should be stateless.**

### How does the framework (work)?
1. `main.go` initializes an empty grid with `HIDDEN` in every element of the `playerBoard`
2. While we have not won yet (validated by the fact there are no more unrevealed cells that aren't `MINE`s)
3. `main.go` asks `solveStep` in `solve.go` to generate a list of cells to reveal (`reveals`) and a list of cells to flag (`flags`), given the `playerBoard`. `reveals` can not be empty, if it is, we fail.
4. `main.go` verifies that every cell in `reveals` is currently hidden, otherwise we fail.
5. For every `revealCoord` in `reveals`:
   1. If it's the first turn, `main.go` takes the first cell in `reveal` and generates a `solution` until the first cell has no adjacent mines.
   2. If `revealCoord` is marked as a `FLAG`, we fail
   3. If `revealCoord` is a `MINE`, we fail
   4. If `revealCoord` is `HIDDEN`, we reveal adjacent cells recursively up to cells that are non-zero.
6. For every `flagCoord` in `flags` toggle the cell's `flag` status unless it is already revealed, then we fail.
