package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartTwoSolution() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	dict := make(map[int]int, 0)     // for numbers from the 2nd column
	nums := make([]int, 0, linesCnt) // for numbers from the 1st column

	for scanner.Scan() {
		line := scanner.Text()
		numsInLine := strings.Split(line, sep)
		num, _ := strconv.Atoi(numsInLine[0])
		nums = append(nums, num)
		num, _ = strconv.Atoi(numsInLine[1])
		dict[num]++
	}

	simScore := 0
	for _, num := range nums {
		simScore += (num * dict[num])
	}

	fmt.Println(simScore)
}
