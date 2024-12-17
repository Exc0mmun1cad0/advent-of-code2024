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

func concatSlices(s1, s2 []int) []int {
	result := make([]int, 0, len(s1)+len(s2))

	for _, el := range s1 {
		result = append(result, el)
	}
	for _, el := range s2 {
		result = append(result, el)
	}

	return result
}
