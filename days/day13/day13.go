package day13

import (
	"fmt"
	"strconv"
	"strings"
)

func Compare(left string, right string) bool {

	return false
}

const (
	OPEN  = -1
	CLOSE = -2
	COMMA = -3
)

func parseNum(a string) int {
	n, _ := strconv.Atoi(a)
	return n
}

func parseLine(line string) []int {
	tokens := []int{}
	curNum := ""

	for _, char := range line {
		switch char {
		case '[':
			if curNum != "" {
				tokens = append(tokens, parseNum(curNum))
				curNum = ""
			}
			tokens = append(tokens, OPEN)
		case ']':
			if curNum != "" {
				tokens = append(tokens, parseNum(curNum))
				curNum = ""
			}
			tokens = append(tokens, CLOSE)
		case ',':
			if curNum != "" {
				tokens = append(tokens, parseNum(curNum))
				curNum = ""
			}
			tokens = append(tokens, COMMA)
		default:
			curNum += string(char)
		}
	}

	return tokens
}

func compare(a []int, b []int) bool {

	ai := 0
	bi := 0

	depthA := 0
	depthB := 0

	for {
		if ai >= len(a) {
			return true
		}

		if bi >= len(b) {
			return false
		}

		aToken := a[ai]
		bToken := b[bi]

		if aToken == OPEN {
			depthA++
		} else if aToken == CLOSE {
			depthA--
		}

		if bToken == OPEN {
			depthB++
		} else if aToken == CLOSE {
			depthB--
		}

		if aToken >= 0 && bToken >= 0 {
			if aToken < bToken {
				return true
			} else if aToken > bToken {
				return false
			}

			ai++
			bi++
		} else if aToken >= 0 {
			bi++
		} else if bToken >= 0 {
			ai++
		} else {
			ai++
			bi++
		}
	}
}

func Part1(input string) int {

	pairs := strings.Split(strings.TrimSpace(input), "\n\n")

	count := 0

	for i, pair := range pairs {
		lines := strings.Split(pair, "\n")
		fmt.Println("Comparing", i, lines)
		result := compare(parseLine(lines[0]), parseLine(lines[1]))

		if result {
			count += i + 1
		}
	}

	return count
}

func Part2(input string) int {
	return 12
}

func Run(input string) {
	fmt.Println("Day 13 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
