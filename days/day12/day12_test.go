package day12

import (
	"testing"
)

const example string = `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

func TestPart1(t *testing.T) {
	t.Log("Testing Day12 Part 1")

	want := 31
	got := Part1(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Log("Testing Day12 Part 2")

	want := 29
	got := Part2(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
