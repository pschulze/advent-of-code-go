package main

import (
	"embed"
	"fmt"
	"strconv"

	"github.com/pschulze/advent-of-code-go/util"
)

//go:embed input.txt
var f embed.FS

func main() {
	data, _ := f.Open("input.txt")
	lines := util.ReadLines(data)

	intLines := intLines(lines)

	fmt.Print("Part 1: ")
	fmt.Println(part1(intLines))

	fmt.Print("Part 2: ")
	fmt.Println(part2(intLines))
}

func intLines(strLines []string) []int {
	var intLines []int
	for _, line := range strLines {
		num, _ := strconv.Atoi(line)
		intLines = append(intLines, num)
	}

	return intLines
}

func part1(values []int) int {
	if len(values) < 2 {
		return 0
	}

	increaseCount := 0
	for i, value := range values {
		if i > 0 && values[i-1] < value {
			increaseCount++
		}
	}

	return increaseCount
}

func part2(values []int) int {
	if len(values) < 3 {
		return 0
	}

	increaseCount := 0
	oldSum := values[0] + values[1] + values[2]

	for i := range values {
		currSum := 0
		if i >= 3 {
			currSum = values[i] + values[i-1] + values[i-2]

			if currSum > oldSum {
				increaseCount++
			}

			oldSum = currSum
		}
	}

	return increaseCount
}
