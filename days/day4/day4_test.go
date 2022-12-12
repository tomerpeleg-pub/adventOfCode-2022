package day4

import "testing"

const example4 string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestDay4Part1(t *testing.T) {
	t.Log("Testing Day4 Part 1")

	want := 2
	got := Day4Part1(example4)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestDay4Part2(t *testing.T) {
	t.Log("Testing Day4 Part 2")

	want := 4
	got := Day4Part2(example4)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
