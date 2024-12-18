package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func PartTwoSolution() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	r := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)

	sum := 0
	flag := true // indicates whether next instruction will be executed or not
	for scanner.Scan() {
		line := scanner.Text()
		instructions := r.FindAllStringSubmatch(line, -1)

		for _, instruction := range instructions {
			switch instruction[0] {
			case `don't()`:
				flag = false
			case `do()`:
				flag = true
			default:
				if flag {
					nums1, _ := strconv.Atoi(instruction[1])
					nums2, _ := strconv.Atoi(instruction[2])
					sum += (nums1 * nums2)
				}
			}
		}
	}

	fmt.Println(sum)
}
