package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func PartOneSolution() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		instructions := r.FindAllStringSubmatch(line, -1)

		for _, instruction := range instructions {
			num1, _ := strconv.Atoi(instruction[1])
			num2, _ := strconv.Atoi(instruction[2])

			sum += (num1 * num2)
		}
	}

	fmt.Println(sum)
}
