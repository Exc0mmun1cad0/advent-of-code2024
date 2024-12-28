package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Recursion for task 2
func rec2(arr []int, pos int, initNum int, sample int) bool {
	if pos == len(arr) {
		return initNum == sample
	}

	return rec2(arr, pos+1, initNum+arr[pos], sample) || rec2(arr, pos+1, initNum*arr[pos], sample) || rec2(arr, pos+1, concat(initNum, arr[pos]), sample)
}

// x ** y
func pow(x, y int) int {
	res := x
	for i := 0; i < y-1; i++ {
		res *= x
	}

	return res
}

// Ooperation of concatenation of 2 numbers.
// Example: 12 || 345 -> 12345
func concat(num1, num2 int) int {
	res := num2
	degree := int(math.Log10(float64(num2))) + 1

	for num1 != 0 {
		res += (num1 % 10) * pow(10, degree)
		num1 /= 10
		degree++
	}

	return res
}

func PartTwoSolution() {
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

		if rec2(nums, 1, nums[0], testVal) {
			result += testVal
		}
	}

	fmt.Println(result)
}
