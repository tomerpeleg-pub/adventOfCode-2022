package days

import "testing"

const example3 string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestDay3Part1(t *testing.T) {
	t.Log("Testing Day3 Part 1")

	want := 157
	got := Day3Part1(example3)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestDay3Part2(t *testing.T) {
	t.Log("Testing Day3 Part 2")

	want := 70
	got := Day3Part2(example3)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
