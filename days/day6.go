package days

import (
	"fmt"
)

func findMarker(input string, n int) int {
	counts := map[rune]int{}
	toRemove := rune(input[0])

main:
	for i, char := range input {
		counts[char]++

		if i < n {
			continue
		}

		counts[toRemove]--

		for _, count := range counts {
			if count > 1 {
				toRemove = rune(input[i-(n-1)])
				continue main
			}
		}

		return i + 1
	}

	return -1
}

func Day6Part1(input string) int {
	return findMarker(input, 4)
}

func Day6Part2(input string) int {
	return findMarker(input, 14)
}

func Day6(input string) {
	fmt.Println("Day 6 -----")
	fmt.Println("Part 1:", Day6Part1(input))
	fmt.Println("Part 2:", Day6Part2(input))
}
