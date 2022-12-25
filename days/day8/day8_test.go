package day8

import (
	"testing"
)

const example8 string = `

30373
25512
65332
33549
35390

`

func TestPart1(t *testing.T) {
	t.Log("Testing Day8 Part 1")

	want := 21
	got := Part1(example8)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Log("Testing Day8 Part 2")

	want := 8
	got := Part2(example8)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
