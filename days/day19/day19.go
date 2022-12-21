package day19

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var types = map[string]int{
	"ore":      0,
	"clay":     1,
	"obsidian": 2,
	"geode":    3,
}

type Robot [4]int

func (r Robot) String() string {
	str := fmt.Sprintf("{ o: %v, c: %v, o: %v, g: %v }", r[0], r[1], r[2], r[3])
	return str
}

type Blueprint [4]Robot

type State struct {
	robots Robot
	stock  Robot
	min    int
}

func parseInput(input string) []Blueprint {
	blueprintRegex := regexp.MustCompile(`Each (\w+) robot costs (\d+) (\w+)( and (\d+) (\w+))?.`)

	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	blueprints := []Blueprint{}

	for scanner.Scan() {
		line := scanner.Text()
		match := blueprintRegex.FindAllStringSubmatch(line, -1)
		blueprint := Blueprint{}

		for _, result := range match {
			robot, cost1, cost1Type, cost2, cost2Type := result[1], result[2], result[3], result[5], result[6]

			c1, _ := strconv.Atoi(cost1)
			blueprint[types[robot]][types[cost1Type]] = c1

			if cost2 != "" {
				c2, _ := strconv.Atoi(cost2)
				blueprint[types[robot]][types[cost2Type]] = c2
			}
		}

		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}

func (s State) Simulate() State {
	s2 := State{stock: s.stock, robots: s.robots, min: s.min}

	for i, r := range s.robots {
		s2.stock[i] += r
	}

	s2.min++
	return s2
}

func (state State) Purchase(blueprint Blueprint, robot int) (State, bool) {
	newState := State{min: state.min}

	for i, cost := range blueprint[robot] {
		if state.stock[i] < cost {
			return state, false
		}

		newState.stock[i] = state.stock[i] - cost
		newState.robots[i] = state.robots[i]
	}

	newState.robots[robot]++

	return newState, true
}

func (state State) Value() float64 {
	s := 0.0

	for i := 0; i < 4; i++ {
		s += float64((1 - 25/(state.min+1)) * state.stock[i] * int(math.Pow10(i)))
		s += float64((25 / (state.min + 1)) * state.robots[i] * int(math.Pow10(i)))
		s += float64(state.min) / 10.0
	}

	return s
}

func (state State) Max(blueprint Blueprint) int {
	n := 23 - state.min

	return ((n / 2) * (n + 1)) + state.stock[3]
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Value() < pq[j].Value()
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}
func (pq *PriorityQueue) Pop() (popped interface{}) {
	popped = (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return
}

func optimise(blueprint Blueprint) int {
	initRobots := Robot{1, 0, 0, 0}
	initStock := Robot{0, 0, 0, 0}
	initState := State{
		robots: initRobots,
		stock:  initStock,
	}

	q := PriorityQueue{}
	q.Push(initState)
	// heap.Init(q)
	// heap.Push(q, initState)

	highest := 0
	highestState := initState

	cache := map[float64]bool{}
	cache[initState.Value()] = true

	h := initState.Value()

	for q.Len() > 0 {
		state := q.Pop().(State)
		state = state.Simulate()

		if state.Value() > h {
			h = state.Value()
			fmt.Println("New value", h, state)
		}
		if state.stock[3] > highest {
			highest = state.stock[3]
			highestState = state
			fmt.Println("New highest", highest, state, state.Value(), q.Len())
		}

		if state.min > 23 {
			continue
		}

		for i := 0; i < 4; i++ {
			newState, ok := state.Purchase(blueprint, i)

			if ok && !cache[newState.Value()] && newState.Max(blueprint) > highest {
				q.Push(newState)
			}
		}

		if !cache[state.Value()] && state.Max(blueprint) > highest {
			q.Push(state)
		}
	}

	return (highestState.stock[3])
}

func Part1(input string) int {
	blueprints := parseInput(input)

	for i, b := range blueprints {
		fmt.Println("==== BLUEPRINT", i, "====", optimise(b))
	}

	return 12
}

func Part2(input string) int {
	return 12
}

func Run(input string) {
	fmt.Println("Day 19 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
