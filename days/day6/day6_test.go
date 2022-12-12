package day6

import (
	"testing"
)

type example struct {
	i string
	e int
}

var example6 []example = []example{
	{i: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`, e: 7},
	{i: `bvwbjplbgvbhsrlpgdmjqwftvncz`, e: 5},
	{i: `nppdvjthqldpwncqszvftbrmjlhg`, e: 6},
	{i: `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`, e: 10},
	{i: `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`, e: 11},
}

func TestDay6Part1(t *testing.T) {
	t.Log("Testing Day6 Part 1")

	for _, ex := range example6 {
		want := ex.e
		got := Day6Part1(ex.i)

		if got != want {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	}
}

var example6_p2 []example = []example{
	{i: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`, e: 19},
	{i: `bvwbjplbgvbhsrlpgdmjqwftvncz`, e: 23},
	{i: `nppdvjthqldpwncqszvftbrmjlhg`, e: 23},
	{i: `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`, e: 29},
	{i: `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`, e: 26},
}

func TestDay6Part2(t *testing.T) {
	t.Log("Testing Day6 Part 2")

	for _, ex := range example6_p2 {
		want := ex.e
		got := Day6Part2(ex.i)

		if got != want {
			t.Errorf("Got %v but wanted %v", got, want)
		}
	}
}
