package main

const (
	inputFile = "input.txt"
	linesCnt  = 1e3
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
