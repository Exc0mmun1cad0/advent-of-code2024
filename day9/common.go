package main

import (
	"bufio"
	"bytes"
	"strconv"

	"github.com/pkg/errors"
)

func ReadLongLine(reader *bufio.Reader) (string, error) {
	const op = "readLongLine"

	var buffer bytes.Buffer

	for line, isPrefix, err := reader.ReadLine(); isPrefix || len(line) != 0; {
		if err != nil {
			return "", errors.Wrap(err, op)
		}

		buffer.Write(line)
		line, isPrefix, err = reader.ReadLine()
	}

	return buffer.String(), nil
}

func RepeatedSlice(n int, num int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, num)
	}

	return res
}

func buildDisk(diskMap string) ([]int, error) {
	const op = "buildDisk"

	var result []int // there will be memory

	var isFree bool     // determines whether next sym in disk map is length of file or free space
	var currFileNum int // number of current file

	for _, sym := range diskMap {
		var toAdd int // symbol which will be repeated currFileNum times and added to memory
		num, err := strconv.Atoi(string(sym))
		if err != nil {
			return nil, errors.Wrap(err, op)
		}

		if !isFree {
			toAdd = currFileNum
			currFileNum++
		} else {
			toAdd = -1
		}

		result = append(result, RepeatedSlice(num, toAdd)...)
		isFree = !isFree
	}

	return result, nil
}

func findLastBusy(memory []int) int {
	for i := len(memory) - 1; i >= 0; i-- {
		if memory[i] != -1 {
			return i
		}
	}

	return -1
}
