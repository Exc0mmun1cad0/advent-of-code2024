package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BFS
func detectTrailheads(y, x int, area [][]int) (int, int) {
	w, h := len(area[0]), len(area)            // width and height of area
	trailheads := make(map[[2]int]struct{}, 0) // how many 9s are reachable from vertex(y, x)
	sum := 0                                   // number of distinct trailheads from vertex(y, x)

	queue := list.New()
	queue.PushBack([]int{y, x})

	for queue.Len() != 0 {

		curr := queue.Front() // first vertex in queue
		queue.Remove(curr)

		coords := curr.Value.([]int) // coordinates of current (first in queue) vertex

		// check if its the end of trailhead
		if area[coords[0]][coords[1]] == 9 {
			trailheads[[2]int{coords[0], coords[1]}] = struct{}{}
			sum++
			continue
		}

		y, x = coords[0], coords[1]
		nexts := [][2]int{{y, x + 1}, {y + 1, x}, {y, x - 1}, {y - 1, x}}
		for _, next := range nexts {
			i, j := next[0], next[1]
			if i >= 0 && j >= 0 && i < h && j < w {
				if area[i][j]-area[y][x] == 1 {
					queue.PushBack([]int{i, j})
				}
			}
		}
	}

	return len(trailheads), sum
}

func main() {
	file, _ := os.Open("data/input.txt")
	scanner := bufio.NewScanner(file)

	area := make([][]int, 0)
	zeros := make([][]int, 0)
	j := 0 // current number line
	for scanner.Scan() {
		line := scanner.Text()

		// converting and adding line to map (area)
		strNums := strings.Split(line, "")
		nums := make([]int, 0, len(strNums))
		for i, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			nums = append(nums, num)
			if num == 0 {
				zeros = append(zeros, []int{j, i})
			}
		}
		area = append(area, nums)
		j++

	}

	sum1, sum2 := 0, 0
	for _, start := range zeros {
		s1, s2 := detectTrailheads(start[0], start[1], area)
		sum1 += s1
		sum2 += s2
	}

	fmt.Printf("1st task: %d", sum1)
	fmt.Printf("2nd task: %d", sum2)
}
