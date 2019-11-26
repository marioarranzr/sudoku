package main

import (
	"fmt"

	"github.com/marioarranzr/sudoku"
)

var easySudoku = []int{
	0, 0, 3, 0, 2, 0, 6, 0, 0,
	9, 0, 0, 3, 0, 5, 0, 0, 1,
	0, 0, 1, 8, 0, 6, 4, 0, 0,

	0, 0, 8, 1, 0, 2, 9, 0, 0,
	7, 0, 0, 0, 0, 0, 0, 0, 8,
	0, 0, 6, 7, 0, 8, 2, 0, 0,

	0, 0, 2, 6, 0, 9, 5, 0, 0,
	8, 0, 0, 2, 0, 3, 0, 0, 9,
	0, 0, 5, 0, 1, 0, 3, 0, 0,
}

var hardSudoku = []int{
	6, 0, 0, 0, 0, 0, 1, 5, 0,
	9, 5, 4, 7, 1, 0, 0, 8, 0,
	0, 0, 0, 5, 0, 2, 6, 0, 0,

	8, 0, 0, 0, 9, 4, 0, 0, 6,
	0, 0, 3, 8, 0, 5, 4, 0, 0,
	4, 0, 0, 3, 7, 0, 0, 0, 8,

	0, 0, 6, 9, 0, 3, 0, 0, 0,
	0, 2, 0, 0, 4, 7, 8, 9, 3,
	0, 4, 9, 0, 0, 0, 0, 0, 5,
}

func main() {
	var (
		e, h []int
		err  error
	)
	e, err = sudoku.Solver(easySudoku)
	if err != nil {
		fmt.Println(err)
	}
	h, err = sudoku.Solver(hardSudoku)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(
		"%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n",
		e[:9],
		e[9:18],
		e[18:27],
		e[27:36],
		e[36:45],
		e[45:54],
		e[54:63],
		e[63:72],
		e[72:81],
		h[:9],
		h[9:18],
		h[18:27],
		h[27:36],
		h[36:45],
		h[45:54],
		h[54:63],
		h[63:72],
		h[72:81],
	)
}
