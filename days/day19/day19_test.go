package day19

import (
	"testing"
)

const example string = `
Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
`

func TestPart1(t *testing.T) {
	t.Log("Testing Day19 Part 1")

	want := 21
	got := Part1(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Log("Testing Day19 Part 2")

	want := 24933642
	got := Part2(example)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
