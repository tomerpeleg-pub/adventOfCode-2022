package day2

import (
	"fmt"
	"strings"
)

var moves map[string]int = map[string]int{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

func scoreRound(p1 int, p2 int) int {
	score := p2

	switch p2 - p1 {
	case 0:
		score += 3
	case 1, -2:
		score += 6
	}

	return score
}

func parseRound(line string) (op int, me int) {
	round := strings.Split(line, "")
	return moves[round[0]], moves[round[2]]
}

func Day2Part1(input string) int {
	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		op, me := parseRound(line)
		total += scoreRound(op, me)
	}

	return total
}

func Day2Part2(input string) int {
	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		op, me := parseRound(line)

		switch me {
		case 1:
			if op == 1 {
				total += scoreRound(op, 3)
			} else {
				total += scoreRound(op, op-1)
			}
		case 2:
			total += scoreRound(op, op)
		case 3:
			if op == 3 {
				total += scoreRound(op, 1)
			} else {
				total += scoreRound(op, op+1)
			}
		}

	}

	return total
}

func Day2(input string) {
	fmt.Println("Day 2 -----")
	fmt.Println("Part 1:", Day2Part1(input))
	fmt.Println("Part 2:", Day2Part2(input))
}
