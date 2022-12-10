package days

import (
	"fmt"
	"strings"
)

type rucksack = map[rune]bool

func parseRucksack(rucksackStr string) rucksack {
	items := rucksack{}

	for _, item := range rucksackStr {
		items[item] = true
	}

	return items
}

func findMatching(rucksacks []rucksack) rune {

outside:
	for key := range rucksacks[0] {
		for _, rucksack := range rucksacks {
			if !rucksack[key] {
				continue outside
			}
		}

		return key
	}

	return 'a'
}

func getItemScore(item rune) int {
	if item <= 'Z' {
		return 27 + (int(item) - 'A')
	} else {
		return 1 + (int(item) - 'a')
	}
}

func Day3Part1(input string) int {
	lines := strings.Split(input, "\n")
	score := 0

	// Loop through each rucksack
	for _, line := range lines {
		lineLen := len(line) / 2
		rucksacks := []rucksack{
			parseRucksack(line[0:lineLen]),
			parseRucksack(line[lineLen:]),
		}
		matchingItem := findMatching(rucksacks)
		score += getItemScore(matchingItem)
	}

	return score
}

func Day3Part2(input string) int {
	lines := strings.Split(input, "\n")
	score := 0

	// Loop through each rucksack
	for i := 0; i < len(lines)-2; i += 3 {
		rucksacks := []rucksack{
			parseRucksack(lines[i]),
			parseRucksack(lines[i+1]),
			parseRucksack(lines[i+2]),
		}
		matchingItem := findMatching(rucksacks)
		score += getItemScore(matchingItem)
	}

	return score
}

func Day3(input string) {
	fmt.Println("Day 3 -----")
	fmt.Println("Part 1:", Day3Part1(input))
	fmt.Println("Part 2:", Day3Part2(input))
}
