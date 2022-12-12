package dayX

import (
	"testing"
)

const example string = `
`

func TestPart1(t *testing.T) {
	t.Log("Testing DayX Part 1")

	want := 21
	got := Part1(example8)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Log("Testing DayX Part 2")

	want := 24933642
	got := Part2(example8)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
