package days

import "testing"

const example string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestDay1Part1(t *testing.T) {
	t.Log("Testing Day1 Part 1")

	want := 24000
	got := Day1Part1(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestDay1Part2(t *testing.T) {
	t.Log("Testing Day1 Part 2")

	want := 45000
	got := Day1Part2(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
