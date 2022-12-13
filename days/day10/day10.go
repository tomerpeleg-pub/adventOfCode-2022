package day10

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	Noop = 0
	Addx = 1
)

type Instruction struct {
	Type int
	Val  int
}

func parseInsturction(input string) Instruction {
	vals := strings.Fields(input)

	switch vals[0] {
	case "noop":
		return Instruction{Type: Noop}
	case "addx":
		num, _ := strconv.Atoi(vals[1])
		return Instruction{Type: Addx, Val: num}
	}

	return Instruction{}
}

func parseInput(input string) []Instruction {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	instructions := []Instruction{}

	// Loop through each line
	for scanner.Scan() {
		instructions = append(instructions, parseInsturction(scanner.Text()))
	}

	return instructions
}

func Part1(input string) int {
	instructions := parseInput(input)

	X := 1
	runningX := 0
	cycle := 0

	for _, instruction := range instructions {
		cycle++
		if cycle%40 == 20 {
			runningX += X * cycle
		}

		if instruction.Type == Addx {
			cycle++
			if cycle%40 == 20 {
				runningX += X * cycle
			}
			X += instruction.Val
		}
	}

	return runningX
}

func Render(line string, X int) string {
	low := (X - 1) % 40
	high := (X + 1) % 40
	pos := len(line)

	if pos >= low && pos <= high {
		line += "#"
	} else {
		line += "."
	}

	return line
}

func Cycle(cycle int, X int, line string, screen string) (int, int, string, string) {
	cycle++
	line = Render(line, X)
	if len(line) == 40 {
		screen += "\n" + line
		line = ""
	}

	return cycle, X, line, screen
}

func Part2(input string) string {
	instructions := parseInput(input)

	X := 1
	cycle := 0
	screen := ""
	line := ""

	for _, instruction := range instructions {
		cycle, X, line, screen = Cycle(cycle, X, line, screen)

		if instruction.Type == Addx {
			cycle, X, line, screen = Cycle(cycle, X, line, screen)
			X += instruction.Val
		}
	}

	return screen
}

func Run(input string) {
	fmt.Println("Day 10 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:\n", Part2(input))
}
