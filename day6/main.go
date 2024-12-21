package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inputFile = "input.txt"
)

// possible directions for movements
const (
	up = iota
	down
	left
	right
)

func main() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var isSet bool
	x, y := -1, -1

	field := make([]string, 0)
	i := 0
	// scanning field and search of start
	for scanner.Scan() {
		line := scanner.Text()

		if !isSet {
			x = strings.Index(line, "^")
			if x != -1 {
				y = i
				isSet = true
			}
		}

		field = append(field, line)
		i++
	}

	fmt.Println(PartOneSolution(x, y, field))
}

func PartOneSolution(x, y int, field []string) int {
	currDir := up                     // current direction. Default is up
	visited := make(map[int]struct{}) // set for storaging positions where guard has been
	visited[y*1000+x] = struct{}{}

	// moving until cursor goes out of field
	for 0 <= x && x < len(field[0]) && 0 <= y && y < len(field) {
		// Define next position
		nextX, nextY := x, y
		switch currDir {
		case up:
			nextY--
		case down:
			nextY++
		case left:
			nextX--
		case right:
			nextX++
		}

		if nextX < 0 || nextX >= len(field[0]) || nextY < 0 || nextY >= len(field) {
			break
		}

		if field[nextY][nextX] == byte('.') || field[nextY][nextX] == byte('^') {
			x, y = nextX, nextY
			visited[y*1000+x] = struct{}{}
		} else if field[nextY][nextX] == byte('#') {
			switch currDir {
			case up:
				currDir = right
			case right:
				currDir = down
			case down:
				currDir = left
			case left:
				currDir = up
			}
		}
	}

	return len(visited)
}
