package days

import (
	"fmt"
	"strconv"
	"strings"
)

type assignment struct {
	Low  int
	High int
	Size int
}

func parseAssignment(str string) assignment {
	vals := strings.Split(str, "-")
	low, _ := strconv.ParseInt(vals[0], 0, 32)
	high, _ := strconv.ParseInt(vals[1], 0, 32)

	return assignment{
		Low:  int(low),
		High: int(high),
		Size: int(high - low),
	}
}

func parseLine(line string) (first assignment, second assignment) {
	elves := strings.Split(line, ",")
	return parseAssignment(elves[0]), parseAssignment(elves[1])
}

func Day4Part1(input string) int {
	fullyContainsCount := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		first, second := parseLine(line)

		if first.Size > second.Size {
			if first.Low <= second.Low && first.High >= second.High {
				fullyContainsCount++
			}
		} else {
			if second.Low <= first.Low && second.High >= first.High {
				fullyContainsCount++
			}
		}
	}

	return fullyContainsCount
}

func Day4Part2(input string) int {
	intersectionCount := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		first, second := parseLine(line)

		if !(first.High < second.Low || first.Low > second.High) {
			intersectionCount++
		}
	}

	return intersectionCount
}

func Day4(input string) {
	fmt.Println("Day 4 -----")
	fmt.Println("Part 1:", Day4Part1(input))
	fmt.Println("Part 2:", Day4Part2(input))
}
