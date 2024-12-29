package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Separation of numbers that have even number of digits, n is number of digits.
// Example: 1234 -> 12 and 34
func separate(num int, n int) []int {
	var num2 int
	power := 1
	for i := 0; i < n/2; i++ {
		num2 += (num % 10) * power
		power *= 10
		num /= 10
	}

	return []int{num, num2}
}

func cntDigits(num int) int {
	return len(strconv.Itoa(num))
}

// Turns old slice of stones to new slice after blink operation
func blink(old map[int]int) map[int]int {
	new := make(map[int]int, 0)

	for key, val := range old {
		if key == 0 {
			new[1] += val
		} else if digits := cntDigits(key); digits%2 == 0 {
			seps := separate(key, digits)
			new[seps[0]] += val
			new[seps[1]] += val
		} else {
			new[key*2024] += val
		}
	}

	return new
}

func task(nums map[int]int, n int) int {
	for i := 0; i < n; i++ {
		nums = blink(nums)
	}
	sum := 0
	for _, val := range nums {
		sum += val
	}

	return sum
}

func main() {
	data, _ := os.ReadFile("data/input.txt")
	strNums := strings.Split(strings.TrimSpace(string(data)), " ")

	nums1 := make(map[int]int, 0)
	nums2 := make(map[int]int, 0)
	for _, strNum := range strNums {
		num, _ := strconv.Atoi(strNum)
		nums1[num]++
		nums2[num]++
	}

	fmt.Printf("1st task: %d\n", task(nums1, 25))
	fmt.Printf("2nd task: %d\n", task(nums2, 75))
}
