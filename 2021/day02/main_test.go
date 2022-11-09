package main

import (
	"reflect"
	"testing"
)

func TestParseCommands(t *testing.T) {
	result := parseCommands([]string{"forward 1", "down 3", "up 2"})

	expectedVecs := []vector{
		{direction: "forward", magnitude: 1},
		{direction: "down", magnitude: 3},
		{direction: "up", magnitude: 2},
	}

	if !reflect.DeepEqual(result, expectedVecs) {
		t.Error("Expected", expectedVecs, " got", result)
	}
}

func TestPart1(t *testing.T) {
	comms := []vector{
		{direction: "forward", magnitude: 1},
		{direction: "down", magnitude: 3},
		{direction: "up", magnitude: 2},
	}

	result := part1(comms)

	if result != 1 {
		t.Error("Expected 1, got", result)
	}
}

func TestPart2(t *testing.T) {
	comms := []vector{
		{direction: "down", magnitude: 4},
		{direction: "up", magnitude: 2},
		{direction: "forward", magnitude: 10},
	}

	// expected hPos: 10, vPos: 10 * (4 - 2) = 20, overall: 10 * 20 = 200
	result := part2(comms)

	if result != 200 {
		t.Error("Expected 200, got", result)
	}
}
