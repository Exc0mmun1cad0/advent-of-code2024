package main

import "fmt"

// Counts how many cells are occupied by file with fileID whose last cell is memory[pos].
// Position is included
func fileSize(memory []int, pos int, fileID int) int {
	res := 0
	for memory[pos] == fileID {
		pos--
		res++
	}

	return res
}

// Returns number of last file in memory and its last cell position.
// If there no files, it will return (-1, -1)
func lastFile(memory []int) (int, int) {
	for i := len(memory) - 1; i >= 0; i-- {
		if memory[i] != -1 {
			return i, memory[i]
		}
	}
	return -1, -1
}

// Scans the memory from the very beginning to find a place to insert a file of length n.
// If place is found, returns its first index. Otherwise, returns -1
func findPlace(memory []int, n int) int {
	curr := 0
	cnt := 0

	for curr < len(memory) {
		if memory[curr] == -1 {
			cnt++
			if cnt == n {
				return curr - n + 1
			}
		} else {
			cnt = 0
		}

		curr++
	}

	return -1
}

func PartTwoSolution(memory []int) {
	lastPos, lastFile := lastFile(memory)

	for lastFile > 0 {
		fileSize := fileSize(memory, lastPos, lastFile)

		placeToSwap := findPlace(memory[:lastPos], fileSize)
		if placeToSwap != -1 {
			// doing swap
			for i := 0; i < fileSize; i++ {
				memory[placeToSwap] = lastFile
				placeToSwap++
				memory[lastPos-i] = -1 // cleaning previous file location
			}
		}

		lastFile--
		for memory[lastPos] != lastFile {
			lastPos--
		}

	}

	// count checksum
	checkSum := 0
	for i, num := range memory {
		if num != -1 {
			checkSum += (i * num)
		}
	}

	fmt.Println(checkSum)
}
