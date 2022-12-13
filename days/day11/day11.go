package day11

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items   []int64
	Op      func(int64) int64
	Test    func(int64) bool
	Tester  int64
	IfTrue  int64
	IfFalse int64
	Count   int64
}

func (monkey Monkey) String() string {
	return fmt.Sprintf("Monkey { Count: %v, Items: %v }", monkey.Count, monkey.Items)
}

func Test(n int64) func(int64) bool {
	return func(old int64) bool {
		return old%n == 0
	}
}

func Op(a int64, b int64, op string) func(int64) int64 {
	return func(old int64) int64 {
		c := a
		d := b
		if a == -1 {
			c = old
		}
		if b == -1 {
			d = old
		}

		if op == "+" {
			return c + d
		}
		return c * d
	}
}

func parseInput(input string) []Monkey {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	monkeys := []Monkey{}
	curMonkey := Monkey{}

	i := 0

	numsReg := regexp.MustCompile(`\d+`)

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()

		switch i % 7 {
		case 1:
			// Starting itmes
			itemsStr := numsReg.FindAllString(line, -1)
			for _, str := range itemsStr {
				k, _ := strconv.ParseInt(str, 10, 64)
				curMonkey.Items = append(curMonkey.Items, int64(k))
			}
		case 2:
			// Operation
			vals := strings.Fields(line)
			a, _ := strconv.ParseInt(vals[3], 10, 64)
			b, _ := strconv.ParseInt(vals[5], 10, 64)
			if vals[3] == "old" {
				a = -1
			}
			if vals[5] == "old" {
				b = -1
			}
			op := vals[4]
			curMonkey.Op = Op(a, b, op)

		case 3:
			// Test
			str := numsReg.FindAllString(line, -1)
			n, _ := strconv.ParseInt(str[0], 10, 64)
			curMonkey.Tester = n
			curMonkey.Test = Test(n)
		case 4:
			// If true
			str := numsReg.FindAllString(line, -1)
			n, _ := strconv.ParseInt(str[0], 10, 64)
			curMonkey.IfTrue = n
		case 5:
			// If false
			str := numsReg.FindAllString(line, -1)
			n, _ := strconv.ParseInt(str[0], 10, 64)
			curMonkey.IfFalse = n
		case 6:
			// End monkey
			monkeys = append(monkeys, curMonkey)
			curMonkey = Monkey{}
		}

		i++
	}
	monkeys = append(monkeys, curMonkey)

	return monkeys
}

func MonkeyBusiness(monkeys []Monkey, rounds int, worryDivider int64) int64 {
	var worryMod int64 = 1
	for _, monkey := range monkeys {
		worryMod *= monkey.Tester
	}

	for i := 0; i < rounds; i++ {
		for m, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkeys[m].Count++
				newWorry := (monkey.Op(item) % worryMod) / worryDivider

				if monkey.Test(newWorry) {
					monkeys[monkey.IfTrue].Items = append(monkeys[monkey.IfTrue].Items, newWorry)
				} else {
					monkeys[monkey.IfFalse].Items = append(monkeys[monkey.IfFalse].Items, newWorry)
				}

				monkeys[m].Items = []int64{}
			}
		}

		if (i+1)%1000 == 0 {
			fmt.Printf("== After round %v ==\n", i+1)

			for m, monkey := range monkeys {
				fmt.Printf("Monkey %v inspected items %v times.\n", m, monkey.Count)
			}
		}
	}

	counts := []int64{}

	for _, monkey := range monkeys {
		counts = append(counts, monkey.Count)
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] < counts[j]
	})

	fmt.Printf("== After round %v ==\n", rounds)
	for m, monkey := range monkeys {
		fmt.Printf("Monkey %v inspected items %v times.\n", m, monkey.Count)
	}

	return counts[len(counts)-1] * counts[len(counts)-2]
}

func Part1(input string) int64 {
	monkeys := parseInput(input)
	return MonkeyBusiness(monkeys, 20, 3)
}

func Part2(input string) int64 {
	monkeys := parseInput(input)
	return MonkeyBusiness(monkeys, 10000, 1)
}

func Run(input string) {
	fmt.Println("Day 11 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
