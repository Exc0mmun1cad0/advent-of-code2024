package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateReport(levels []int) bool {
	var isIncreasing bool

	// check first two levels in report in order to identify trend (increasing or decreasing)
	if levels[0] == levels[1] || abs(levels[0]-levels[1]) < 1 || abs(levels[0]-levels[1]) > 3 {
		return false
	} else if levels[0] < levels[1] {
		isIncreasing = true
	}

	// then check the remaining levels
	for i := 1; i < len(levels)-1; i++ {
		if levels[i] == levels[i+1] {
			return false
		}

		isInc := levels[i] < levels[i+1]
		if isInc != isIncreasing || abs(levels[i]-levels[i+1]) < 1 || abs(levels[i]-levels[i+1]) > 3 {
			return false
		}
	}

	return true
}

func PartOneSolution() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	reports := 0
	for scanner.Scan() {
		line := scanner.Text()
		strNums := strings.Split(line, " ")

		if len(strNums) < 2 { // just in case handle corner case with one (or even less) levels in report
			reports++
			continue
		}

		nums := make([]int, 0, len(strNums))
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			nums = append(nums, num)
		}

		if validateReport(nums) {
			reports++
		}
	}

	fmt.Println(reports)
}
