.PHONY: run
run:
	go run cmd/sudoku/main.go

test:
	go test ./... -v -count=1