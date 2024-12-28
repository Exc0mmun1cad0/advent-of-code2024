package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func recursion(arr []int, pos int, initNum int, sample int) bool {
	if pos == len(arr) {
		return initNum == sample
	}

	return recursion(arr, pos+1, initNum+arr[pos], sample) || recursion(arr, pos+1, initNum*arr[pos], sample)
}

func PartOneSolution() {
	file, _ := os.Open("data/input.txt")
	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		delIndex := strings.Index(line, ":") // index of delimeter ":" between test value and numbers

		testVal, _ := strconv.Atoi(line[:delIndex])
		strNums := strings.Split(line[delIndex+2:], " ") // +2 in order to get rid of colon and first space

		nums := make([]int, 0, len(strNums))
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			nums = append(nums, num)
		}

		if recursion(nums, 1, nums[0], testVal) {
			result += testVal
		}
	}

	fmt.Println(result)
}
