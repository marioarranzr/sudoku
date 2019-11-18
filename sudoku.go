package sudoku

import "fmt"

// Solver is the main task of this app
func Solver(puzzle []int) ([]int, error) {
	if len(puzzle) != 81 {
		return nil, fmt.Errorf("the sudoku has to have 9 rows and 9 columns")
	}
	solved := backtrack(puzzle)
	if !solved {
		return nil, fmt.Errorf("the sudoku was impossible to solve")
	}
	return puzzle, nil
}

func backtrack(puzzle []int) bool {
	if isSolved(puzzle) {
		return true
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			pos := (i-1)*9 + j - 1
			if puzzle[pos] == 0 {
				for t := 1; t <= 9; t++ {
					puzzle[pos] = t
					if isValid(puzzle) {
						if backtrack(puzzle) {
							return true
						}
						puzzle[pos] = 0
					} else {
						puzzle[pos] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func isValid(puzzle []int) bool {
	// number of duplicates by row
	for i := 1; i <= 9; i++ {
		counterRow := make([]int, 10)
		for j := 1; j <= 9; j++ {
			posRow := (i-1)*9 + j - 1
			counterRow[puzzle[posRow]]++
		}
		if hasDuplicates(counterRow) {
			return false
		}
	}

	// number of duplicates by column
	for i := 1; i <= 9; i++ {
		counterCol := make([]int, 10)
		for j := 1; j <= 9; j++ {
			posCol := (j-1)*9 + i - 1
			counterCol[puzzle[posCol]]++
		}
		if hasDuplicates(counterCol) {
			return false
		}
	}

	// number of duplicates by section
	for i := 1; i <= 9; i += 3 {
		for j := 1; j <= 9; j += 3 {
			counter := make([]int, 10)
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					pos := (row-1)*9 + col - 1
					counter[puzzle[pos]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter []int) bool {
	for i := 1; i < len(counter); i++ {
		if counter[i] > 1 {
			return true
		}
	}
	return false
}

func isSolved(puzzle []int) bool {
	for i := 0; i < len(puzzle); i++ {
		if puzzle[i] == 0 {
			return false
		}
	}
	return true
}
