package main

import "testing"

func TestPart1Basic(t *testing.T) {
	result := part1([]int{1, 2, 3})
	if result != 2 {
		t.Error("Expected 2, got", result)
	}
}

func TestPart1AllEqual(t *testing.T) {
	result := part1([]int{1, 1, 1})
	if result != 0 {
		t.Error("Expected 0, got", result)
	}
}

func TestPart1Decreasing(t *testing.T) {
	result := part1([]int{3, 2, 1})
	if result != 0 {
		t.Error("Expected 0, got", result)
	}
}

func TestPart1Empty(t *testing.T) {
	result := part1([]int{})
	if result != 0 {
		t.Error("Expected 0, got", result)
	}
}

func TestPart1Mixed(t *testing.T) {
	result := part1([]int{1, 2, 1})
	if result != 1 {
		t.Error("Expected 1, got", result)
	}
}

func TestPart12imple(t *testing.T) {
	result := part2([]int{1, 2, 3, 4})
	if result != 1 {
		t.Error("Expected 1, got", result)
	}
}

func TestPart2AllEqual(t *testing.T) {
	result := part2([]int{1, 1, 1, 1})
	if result != 0 {
		t.Error("Expected 0, got", result)
	}
}

func TestPart2Decreasing(t *testing.T) {
	result := part2([]int{4, 3, 2, 1})
	if result != 0 {
		t.Error("Expected 0, got", result)
	}
}

func TestPart2Mixed(t *testing.T) {
	result := part2([]int{2, 2, 3, 5, 1, 10})
	if result != 2 {
		t.Error("Expected 2, got", result)
	}
}

func TestPart2Empty(t *testing.T) {
	result := part2([]int{})
	if result != 0 {
		t.Error("Expected 0, got", result)
	}
}
