package day25

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

func snifuToInt(snifu string) int {
	nums := make([]int, len(snifu))
	tot := 0

	for i, char := range snifu {
		switch char {
		case '2':
			nums[i] = 2
		case '1':
			nums[i] = 1
		case '0':
			nums[i] = 0
		case '-':
			nums[i] = -1
		case '=':
			nums[i] = -2
		}
	}

	for i, v := range nums {
		k := len(nums) - i - 1
		f := math.Pow(5, float64(k))
		tot += int(f) * v
	}
	return tot
}

func numToSnifu(num int) string {
	if num == 0 {
		return ""
	}

	a := math.Floor(float64(num+2) / 5.0)
	b := (num + 2) % 5

	nums := []string{"=", "-", "0", "1", "2"}
	return numToSnifu(int(a)) + nums[b]
}

func parseInput(input string) []int {

	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)
	results := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		results = append(results, snifuToInt(line))
	}
	return results
}

func Part1(input string) string {
	result := parseInput(input)

	s := 0

	for _, v := range result {
		s += v
	}

	return numToSnifu(s)
}

func Part2(input string) int {
	return 12
}

func Run(input string) {
	fmt.Println("Day 25 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
