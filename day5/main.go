package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input.txt"
)

func handleUpd(nums []int, order map[int]map[int]struct{}) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if _, ok := order[nums[i]][nums[j]]; !ok {
				return false
			}
		}
	}
	return true
}

func checkAndSortUpd(nums []int, order map[int]map[int]struct{}) (bool, []int) {
	isCorrect := true
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if _, ok := order[nums[i]][nums[j]]; !ok {
				nums[i], nums[j] = nums[j], nums[i]
				isCorrect = false
			}
		}
	}

	return isCorrect, nums
}

func main() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rules := make(map[int]map[int]struct{})

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break // case when it's "empty" line
		}

		nums := strings.Split(strings.TrimSpace(line), "|")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		if rules[num1] == nil {
			rules[num1] = make(map[int]struct{})
		}
		rules[num1][num2] = struct{}{}
	}

	// fmt.Println(PartOneSolution(scanner, rules))
	// fmt.Println(PartTwoSolution(scanner, rules))
}

func PartTwoSolution(scanner *bufio.Scanner, rules map[int]map[int]struct{}) int {
	sum := 0 // result sum of middle elements

	for scanner.Scan() {
		line := scanner.Text()

		strNums := strings.Split(line, ",")
		nums := make([]int, 0, len(strNums))
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			nums = append(nums, num)
		}

		if isCorrect, sortedNums := checkAndSortUpd(nums, rules); !isCorrect {
			sum += sortedNums[len(sortedNums)/2]
		}
	}

	return sum
}

func PartOneSolution(scanner *bufio.Scanner, rules map[int]map[int]struct{}) int {
	sum := 0 // result sum of middle elements

	for scanner.Scan() {
		line := scanner.Text()

		strNums := strings.Split(line, ",")
		nums := make([]int, 0, len(strNums))
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			nums = append(nums, num)
		}

		if handleUpd(nums, rules) {
			sum += nums[len(nums)/2]
		}
	}

	return sum
}
