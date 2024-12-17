package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func PartOneSolution() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	nums1, nums2 := make([]int, 0, linesCnt), make([]int, 0, linesCnt)

	for scanner.Scan() {
		line := scanner.Text()
		numsInLine := strings.Split(line, sep) // always 2 nums in line

		num, _ := strconv.Atoi(numsInLine[0])
		nums1 = append(nums1, num)
		num, _ = strconv.Atoi(numsInLine[1])
		nums2 = append(nums2, num)
	}

	// Sorting nums in both slices
	sort.Ints(nums1)
	sort.Ints(nums2)
	// Make sure that they are of the same length
	if len(nums1) != len(nums2) {
		log.Fatal("Slices are of different size. Something went wrong...")
	}

	totalDist := 0
	for i := 0; i < len(nums1); i++ {
		totalDist += (abs(nums2[i] - nums1[i]))
	}

	fmt.Println(totalDist)
}
