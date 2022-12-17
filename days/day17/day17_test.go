package day17

import (
	"testing"
)

const example string = `
>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
`

func TestPart1(t *testing.T) {
	t.Log("Testing Day17 Part 1")

	want := 3068
	got := Part1(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Log("Testing Day17 Part 2")

	want := 1514285714288
	got := Part2(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
