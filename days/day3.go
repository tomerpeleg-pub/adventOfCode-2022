package days

import (
	"fmt"
	"strings"
)

func Day3Part1(input string) int {
	lines := strings.Split(input, "\n")
	score := 0

	// Loop through each rucksack
	for _, line := range lines {
		items := map[rune]int{}
		lineLen := len(line) / 2

		// Loop through each item in the rucksack
		for i, char := range line {

			if i < lineLen {
				// First compartment
				// Store each type of item seen (don't care how many)
				items[char] = 1
			} else if items[char] == 1 {
				// Second compartment, only care about items already
				//  seen in the first compartment
				// Mark item as seen in both
				items[char] = 2

				// Add score depending on type of item
				if char <= 'Z' {
					score += 27 + (int(char) - 'A')
				} else {
					score += 1 + (int(char) - 'a')
				}
			}
		}
	}

	return score
}

func Day3(input string) {
	fmt.Println("Day 2 -----")
	fmt.Println("Part 1:", Day3Part1(input))
	// fmt.Println("Part 2:", Day2Part2(input))
}
