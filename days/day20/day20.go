package day20

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	nums := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		nums = append(nums, val)
	}

	return nums
}

func Part1(input string) int {
	vals := parseInput(input)
	l := len(vals)

	dec := make([]int, l)

	for i, v := range vals {
		j := v
		if v < 0 {
			j = l - v
		}
		dec[(i+j)%l] = v
	}
	fmt.Println(dec)

	return 12
}

func Part2(input string) int {
	return 12
}

func Run(input string) {
	fmt.Println("Day 20 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
