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

		for i := 0; i < len(nums); i++ {
			if validateReport(concatSlices(nums[:i], nums[i+1:])) {
				reports++
				break
			}
		}
	}

	fmt.Println(reports)

}
