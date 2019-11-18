package sudoku

import (
	"testing"
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

var easySolution = []int{
	4, 8, 3, 9, 2, 1, 6, 5, 7,
	9, 6, 7, 3, 4, 5, 8, 2, 1,
	2, 5, 1, 8, 7, 6, 4, 9, 3,
	5, 4, 8, 1, 3, 2, 9, 7, 6,
	7, 2, 9, 5, 6, 4, 1, 3, 8,
	1, 3, 6, 7, 9, 8, 2, 4, 5,
	3, 7, 2, 6, 8, 9, 5, 1, 4,
	8, 1, 4, 2, 5, 3, 7, 6, 9,
	6, 9, 5, 4, 1, 7, 3, 8, 2,
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

var hardSolution = []int{
	6, 3, 2, 4, 8, 9, 1, 5, 7,
	9, 5, 4, 7, 1, 6, 3, 8, 2,
	1, 7, 8, 5, 3, 2, 6, 4, 9,
	8, 1, 7, 2, 9, 4, 5, 3, 6,
	2, 9, 3, 8, 6, 5, 4, 7, 1,
	4, 6, 5, 3, 7, 1, 9, 2, 8,
	7, 8, 6, 9, 5, 3, 2, 1, 4,
	5, 2, 1, 6, 4, 7, 8, 9, 3,
	3, 4, 9, 1, 2, 8, 7, 6, 5,
}

var impossibleToSolveSudoku = []int{
	9, 0, 3, 0, 2, 0, 6, 0, 0,
	9, 0, 0, 3, 0, 5, 0, 0, 1,
	0, 0, 1, 8, 0, 6, 4, 0, 0,

	0, 0, 8, 1, 0, 2, 9, 0, 0,
	7, 0, 0, 0, 0, 0, 0, 0, 8,
	0, 0, 6, 7, 0, 8, 2, 0, 0,

	0, 0, 2, 6, 0, 9, 5, 0, 0,
	8, 0, 0, 2, 0, 3, 0, 0, 9,
	0, 0, 5, 0, 1, 0, 3, 0, 0,
}

var errSudoku = []int{
	6,
}

func TestSolver(t *testing.T) {
	type args struct {
		puzzle []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "easySudoku",
			args: args{
				easySudoku,
			},
			want:    easySolution,
			wantErr: false,
		},
		{
			name: "hardSudoku",
			args: args{
				hardSudoku,
			},
			want:    hardSolution,
			wantErr: false,
		},
		{
			name: "notSolvedSudoku",
			args: args{
				impossibleToSolveSudoku,
			},
			wantErr: true,
		},
		{
			name: "errSudoku",
			args: args{
				errSudoku,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solver(tt.args.puzzle)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !compare(got, tt.want) {
				t.Errorf("Solver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compare(puzzle []int, completePuzzle []int) bool {
	for i, n := range puzzle {
		if completePuzzle[i] != n {
			return false
		}
	}
	return true
}
