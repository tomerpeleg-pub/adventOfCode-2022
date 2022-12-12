package day7

import (
	"testing"
)

const example7 string = `
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

func TestDay7Part1(t *testing.T) {
	t.Log("Testing Day7 Part 1")

	want := 95437
	got := Day7Part1(example7)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}

func TestDay7Part2(t *testing.T) {
	t.Log("Testing Day7 Part 2")

	want := 24933642
	got := Day7Part2(example7)

	if got != want {
		t.Errorf("Got %v but wanted %v", got, want)
	}
}
