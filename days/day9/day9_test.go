package day9

import (
	"testing"
)

const example string = `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

func TestPart1(t *testing.T) {
	t.Log("Testing DayX Part 1")

	want := 13
	got := Part1(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

const example2 string = `
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

func TestPart2(t *testing.T) {
	t.Log("Testing DayX Part 2")

	want := 36
	got := Part2(example2)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
