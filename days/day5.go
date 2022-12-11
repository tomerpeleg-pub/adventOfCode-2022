package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type move struct {
	Num  uint64
	From uint64
	To   uint64
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func parseStacks(input string) []stack {
	lines := strings.Split(input, "\n")
	numStacks := len(strings.Fields(lines[len(lines)-1]))

	stacks := make([]stack, numStacks)

	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]

		for k := 0; k < numStacks; k++ {
			char := string(line[(k*4)+1])
			if char != " " {
				stacks[k] = append(stacks[k], char)
			}
		}
	}

	return stacks
}

func parseMoves(input string) []move {
	reg := regexp.MustCompile("[0-9]+")

	moves := []move{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		nums := reg.FindAllString(line, -1)

		num, _ := strconv.ParseUint(nums[0], 10, 32)
		from, _ := strconv.ParseUint(nums[1], 10, 32)
		to, _ := strconv.ParseUint(nums[2], 10, 32)

		m := move{
			Num:  num,
			From: from - 1,
			To:   to - 1,
		}

		moves = append(moves, m)
	}

	return moves
}

func parseInput(input string) ([]stack, []move) {
	parts := strings.Split(input, "\n\n")
	stacks := parseStacks(parts[0])
	moves := parseMoves(parts[1])

	return stacks, moves
}

func Day5Part1(input string) string {
	stacks, moves := parseInput(input)

	for _, move := range moves {
		for i := 0; i < int(move.Num); i++ {
			s, last := stacks[move.From].Pop()
			stacks[move.From] = s
			stacks[move.To] = stacks[move.To].Push(last)
		}
	}

	result := ""

	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}
	return result
}

func Day5Part2(input string) string {
	stacks, moves := parseInput(input)

	for _, move := range moves {
		toMove := stacks[move.From][len(stacks[move.From])-int(move.Num):]
		stacks[move.From] = stacks[move.From][:len(stacks[move.From])-int(move.Num)]
		stacks[move.To] = append(stacks[move.To], toMove...)
	}

	result := ""

	for _, stack := range stacks {
		result += stack[len(stack)-1]
	}
	return result
}

func Day5(input string) {
	fmt.Println("Day 5 -----")
	fmt.Println("Part 1:", Day5Part1(input))
	fmt.Println("Part 2:", Day5Part2(input))
}
