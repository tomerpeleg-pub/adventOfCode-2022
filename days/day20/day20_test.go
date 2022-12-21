package day20

import (
	"testing"
)

const example string = `
1
2
-3
3
-2
0
4
`

func TestPart1(t *testing.T) {
	t.Log("Testing Day20 Part 1")

	want := 3
	got := Part1(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Log("Testing Day20 Part 2")

	want := 1623178306
	got := Part2(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
