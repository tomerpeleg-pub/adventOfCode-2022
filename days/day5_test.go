package days

import "testing"

const example5 string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestDay5Part1(t *testing.T) {
	t.Log("Testing Day5 Part 1")

	want := "CMZ"
	got := Day5Part1(example5)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestDay5Part2(t *testing.T) {
	t.Log("Testing Day5 Part 2")

	want := "MCD"
	got := Day5Part2(example5)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
