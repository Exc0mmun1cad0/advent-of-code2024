package main

import (
	"fmt"
	"slices"
)

func PartOneSolution(memory []int) {
	firstFree := slices.Index(memory, -1)
	lastBusy := findLastBusy(memory)

	for firstFree < lastBusy {
		memory[firstFree], memory[lastBusy] = memory[lastBusy], memory[firstFree] // swap

		// moving firstFree pointer to next free memory cell
		for memory[firstFree] != -1 {
			firstFree++
		}

		// mvoing lastBusy pointer to last busy memory cell
		for memory[lastBusy] == -1 {
			lastBusy--
		}
	}

	sum := 0
	for i, num := range memory[:firstFree] {
		sum += (i * num)
	}

	fmt.Println(sum)
}
