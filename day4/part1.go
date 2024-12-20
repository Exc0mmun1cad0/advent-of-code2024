package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	inputFile = "input.txt"
)

func PartOneSolution() {
	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)
	linesCnt := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		linesCnt++
	}
	lineLength := len(lines[0])

	cnt := 0

	// Firstly, going horizontally
	for _, line := range lines {
		i := 0
		for i < len(line)-3 {
			if line[i:i+4] == "XMAS" || line[i:i+4] == "SAMX" {
				cnt++
				i += 3
			} else {
				i += 1
			}
		}
	}

	// Going vertically
	for i := 0; i < lineLength; i++ {
		j := 0
		var b bytes.Buffer
		for j < linesCnt-3 {
			for pos := j; pos < j+4; pos++ {
				b.WriteByte(lines[pos][i])
			}
			if currWord := b.String(); currWord == "XMAS" || currWord == "SAMX" {
				cnt++
				j += 3
			} else {
				j++
			}
			b.Reset()
		}
	}

	// Diagonally
	for i := 0; i < linesCnt; i++ {
		for j := 0; j < lineLength; j++ {
			var b bytes.Buffer
			// Diagonally right down from current position
			if i < linesCnt-3 && j < lineLength-3 {
				for pos := 0; pos < 4; pos++ {
					b.WriteByte(lines[i+pos][j+pos])
				}
				if currWord := b.String(); currWord == "XMAS" || currWord == "SAMX" {
					cnt++
				}
				b.Reset()
			}

			// ... left up ...
			if i >= 3 && j < lineLength-3 {
				for pos := 0; pos < 4; pos++ {
					b.WriteByte(lines[i-pos][j+pos])
				}
				if currWord := b.String(); currWord == "XMAS" || currWord == "SAMX" {
					cnt++
				}
				b.Reset()
			}
		}
	}

	fmt.Println(cnt)
}
