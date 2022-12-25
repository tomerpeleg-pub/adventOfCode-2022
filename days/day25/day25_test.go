package day25

import (
	"testing"
)

const example string = `
1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
`

func TestPart1(t *testing.T) {
	t.Log("Testing Day25 Part 1")

	want := "2=-1=0"
	got := Part1(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Log("Testing Day25 Part 2")

	want := 24933642
	got := Part2(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
