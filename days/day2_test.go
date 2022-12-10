package days

import "testing"

const example2 string = `A Y
B X
C Z`

func TestDay2Part1(t *testing.T) {
	t.Log("Testing Day2 Part 1")

	want := 15
	got := Day2Part1(example2)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestDay2Part12(t *testing.T) {
	t.Log("Testing Day2 Part 2")

	want := 12
	got := Day2Part2(example2)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
