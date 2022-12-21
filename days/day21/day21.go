package day21

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	Name   string
	Shout  int
	Wait   []string
	Action string
	Done   bool
}

type Monkeys map[string]Monkey

func parseInput(input string) map[string]Monkey {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	monkeys := Monkeys{}

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Fields(line)

		monkey := Monkey{Name: vals[0][0:4]}

		if len(vals) == 2 {
			val, _ := strconv.Atoi(vals[1])
			monkey.Shout = val
			monkey.Done = true
		} else {
			monkey.Wait = []string{vals[1], vals[3]}
			monkey.Action = vals[2]
			monkey.Done = false
		}

		monkeys[monkey.Name] = monkey
	}

	return monkeys
}

func GetShout(monkeys Monkeys, id string) int {
	monkey := monkeys[id]
	if monkey.Done {
		return monkey.Shout
	}

	vals := [2]int{}
	for i, id2 := range monkey.Wait {
		vals[i] = GetShout(monkeys, id2)
	}

	tot := 0
	switch monkey.Action {
	case "+":
		tot = vals[0] + vals[1]
	case "-":
		tot = vals[0] - vals[1]
	case "*":
		tot = vals[0] * vals[1]
	case "/":
		tot = vals[0] / vals[1]
	}

	monkey.Shout = tot
	monkey.Done = true
	return tot
}

func GetEquasion(monkeys Monkeys, id string) string {
	monkey := monkeys[id]
	if id == "humn" {
		return "x"
	}
	if monkey.Done {
		return fmt.Sprint(monkey.Shout)
	}
	monkey2 := monkeys[id]

	act := monkey2.Action
	if id == "root" {
		act = "="
	}

	v1 := GetEquasion(monkeys, monkey2.Wait[0])
	v2 := GetEquasion(monkeys, monkey2.Wait[1])

	if !strings.Contains(v1, "x") {
		v1 = fmt.Sprint(GetShout(monkeys, monkey2.Wait[0]))
	}
	if !strings.Contains(v2, "x") {
		v2 = fmt.Sprint(GetShout(monkeys, monkey2.Wait[1]))
	}

	return fmt.Sprintf("(%v %v %v)", v1, act, v2)
}

func Part1(input string) int {
	monkeys := parseInput(input)

	return GetShout(monkeys, "root")
}

func Part2(input string) int {
	monkeys := parseInput(input)

	fmt.Println("root equasion:", GetEquasion(monkeys, "root"))

	// stuck that in
	// https://www.mathpapa.com/equation-solver/
	// and you get
	return 3403989691757
}

func Run(input string) {
	fmt.Println("Day 21 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
