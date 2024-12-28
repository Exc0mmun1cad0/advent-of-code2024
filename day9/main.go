package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	reader := bufio.NewReader(f)

	diskMap, err := ReadLongLine(reader)
	_ = err

	memory, _ := buildDisk(diskMap)

	PartOneSolution(memory)
}
