package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/pschulze/advent-of-code-go/util"
)

type vector struct {
	direction string
	magnitude int
}

//go:embed input.txt
var f embed.FS

func main() {
	data, _ := f.Open("input.txt")
	lines := util.ReadLines(data)

	commands := parseCommands(lines)

	fmt.Print("Part 1: ")
	fmt.Println(part1(commands))

	fmt.Print("Part 2: ")
	fmt.Println(part2(commands))
}

func parseCommands(lines []string) []vector {
	var comms []vector
	for _, line := range lines {
		fields := strings.Fields(line)
		direction := fields[0]
		magnitude, _ := strconv.Atoi(fields[1])

		vec := vector{direction: direction, magnitude: magnitude}
		comms = append(comms, vec)
	}

	return comms
}

func part1(commands []vector) int {
	verticalPosition, horizontalPosition := 0, 0

	for _, command := range commands {
		switch command.direction {
		case "forward":
			horizontalPosition += command.magnitude
		case "up":
			verticalPosition -= command.magnitude
		case "down":
			verticalPosition += command.magnitude
		}
	}

	return horizontalPosition * verticalPosition
}

func part2(commands []vector) int {
	verticalPosition, horizontalPosition, aim := 0, 0, 0

	for _, command := range commands {
		switch command.direction {
		case "forward":
			horizontalPosition += command.magnitude
			verticalPosition += aim * command.magnitude
			// if verticalPosition < 0 {
			// 	verticalPosition = 0
			// }
		case "up":
			aim -= command.magnitude
		case "down":
			aim += command.magnitude
		}
	}

	return horizontalPosition * verticalPosition
}
