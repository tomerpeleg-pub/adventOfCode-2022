package day13

import (
	"fmt"
	"sort"
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

	INT = 0
	ARR = 1
)

type Val struct {
	Type int
	Val  int
	Vals []Val
}

func (val Val) String() string {
	str := ""

	if val.Type == ARR {
		str += "["

		for j, v := range val.Vals {
			str += v.String()

			if j < len(val.Vals)-1 {
				str += ","
			}
		}

		str += "]"
	} else {
		str += fmt.Sprint(val.Val)
	}

	return str
}

func parseNum(a string) int {
	n, _ := strconv.Atoi(a)
	return n
}

func scanLine(line string) []int {
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

func parseArray(tokens []int) (Val, int) {
	arr := Val{
		Type: ARR,
		Vals: []Val{},
	}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch token {
		case OPEN:
			result, j := parseArray(tokens[i+1:])
			arr.Vals = append(arr.Vals, result)
			i += j
		case CLOSE:
			return arr, i + 1
		case COMMA:
			continue
		default:
			arr.Vals = append(arr.Vals, Val{Val: token, Type: INT})
		}
	}

	return arr, -1
}

func parseLine(line string) Val {
	tokens := scanLine(line)
	result, _ := parseArray(tokens[1:])

	return result
}

func compare(valA Val, valB Val) (result bool, stop bool) {
	result = true
	stop = false

	if valA.Type == INT && valB.Type == INT {
		if valA.Val < valB.Val {
			return true, true
		} else if valA.Val > valB.Val {
			return false, true
		} else {
			return true, false
		}
	}

	for i := 0; i < len(valA.Vals) || i < len(valB.Vals); i++ {
		if i >= len(valA.Vals) && i < len(valB.Vals) {
			return true, true
		} else if i < len(valA.Vals) && i >= len(valB.Vals) {
			return false, true
		}

		a := valA.Vals[i]
		b := valB.Vals[i]

		if a.Type != b.Type {
			c := Val{}
			if a.Type == INT {
				c.Type = ARR
				c.Vals = []Val{a}
				result, stop = compare(c, b)

				if stop {
					return
				}
			} else {
				c.Type = ARR
				c.Vals = []Val{b}
				result, stop = compare(a, c)

				if stop {
					return
				}
			}
		} else {
			result, stop = compare(a, b)

			if stop {
				return
			}
		}
	}

	return
}

func Part1(input string) int {

	pairs := strings.Split(strings.TrimSpace(input), "\n\n")

	count := 0

	for i, pair := range pairs {
		lines := strings.Split(pair, "\n")
		// fmt.Println("Comparing", i, lines)
		line1 := parseLine(lines[0])
		line2 := parseLine(lines[1])
		result, _ := compare(line1, line2)

		if result {
			count += i + 1
		}

		// fmt.Println("line1: ", line1)
		// fmt.Println("line2: ", line2)
		// fmt.Println("Comparison: ", result)
	}

	return count
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	packet1 := parseLine("[[2]]")
	packet2 := parseLine("[[6]]")

	vals := []Val{
		packet1,
		packet2,
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		val := parseLine(line)
		vals = append(vals, val)
	}

	sort.Slice(vals, func(i, j int) bool {
		result, _ := compare(vals[i], vals[j])
		return result
	})

	p1i, p2i := 0, 0

	for i, val := range vals {
		if val.String() == packet1.String() {
			p1i = i + 1
		} else if val.String() == packet2.String() {
			p2i = i + 1
		}
	}

	return p1i * p2i
}

func Run(input string) {
	fmt.Println("Day 13 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
