package day13

import (
	"testing"
)

const example string = `
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

func TestPart1(t *testing.T) {
	t.Log("Testing Day13 Part 1")

	want := 13
	got := Part1(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Log("Testing Day13 Part 2")

	want := 140
	got := Part2(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
