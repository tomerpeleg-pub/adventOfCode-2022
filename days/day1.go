package days

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tomerpeleg-pub/aoc2022/util"
)

func Day1Part1(input string) int {
	lines := strings.Split(input, "\n")

	current, highest := 0, 0

	for _, line := range lines {
		if line == "" {
			if current > highest {
				highest = current
			}

			current = 0
			continue
		}

		val, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println(err)
			continue
		}

		current += val
	}
	if current > highest {
		highest = current
	}

	return highest
}

func Day1Part2(input string) int {
	lines := strings.Split(input, "\n")

	elfCalories := []int{}
	current := 0

	for _, line := range lines {
		if line == "" {
			elfCalories = append(elfCalories, current)
			current = 0
			continue
		}

		calories, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println(err)
			continue
		}

		current += calories
	}
	elfCalories = append(elfCalories, current)
	sort.Ints(elfCalories)

	highestThree := elfCalories[len(elfCalories)-3:]

	return util.SumInt(highestThree)
}

func Day1(input string) {
	fmt.Println("Day 1 -----")
	fmt.Println("Part 1:", Day1Part1(input))
	fmt.Println("Part 2:", Day1Part2(input))
}
