package day16

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type Valve struct {
	id      string
	flow    int
	tunnels map[string]int
	open    bool
}

func (v Valve) String() string {
	return fmt.Sprintf("%v{ flow: %v, tunnels: %v }", v.id, v.flow, v.tunnels)
}

var valveRegex = regexp.MustCompile(`Valve (?P<id>\w+) has flow rate=(?P<flow>\d+); tunnels? leads? to valves? (?P<tunnels>.*)`)

func parseLine(line string) Valve {
	result := make(map[string]string)
	match := valveRegex.FindStringSubmatch(line)
	for i, name := range valveRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	id := result["id"]
	flow, _ := strconv.Atoi(result["flow"])
	valves := strings.Split(result["tunnels"], ", ")
	tunnels := map[string]int{}

	for _, v := range valves {
		tunnels[v] = 1
	}

	return Valve{
		id, flow, tunnels, false,
	}
}

func parseInput(input string) map[string]Valve {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	valves := map[string]Valve{}

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		valve := parseLine(line)
		valves[valve.id] = valve
	}

	return valves
}

func copyMap(a map[string]bool) map[string]bool {
	b := map[string]bool{}
	for k, v := range a {
		b[k] = v
	}
	return b
}

func bestPath(valves map[string]Valve, opened map[string]bool, start Valve, mins int, total int) int {
	if mins <= 1 {
		return total
	}

	time := mins
	newOpened := copyMap(opened)
	curTotal := total

	if !opened[start.id] && start.flow > 0 {
		time--
		newOpened[start.id] = true
		curTotal += time * start.flow
	}

	curHighest := curTotal

	for id, dist := range start.tunnels {
		if dist > time {
			continue
		}
		if newOpened[id] {
			continue
		}
		v := valves[id]

		result := bestPath(valves, newOpened, v, time-dist, curTotal)

		if result > curHighest {
			curHighest = result
		}
	}

	return curHighest
}

type QueueNode struct {
	id   string
	dist int
}

func getVertexes(valves map[string]Valve, start string) Valve {
	q := stack.New()
	q.Push(QueueNode{start, 0})

	visited := map[string]int{}

	for i := 0; q.Len() > 0; i++ {
		qn := q.Pop().(QueueNode)
		id := qn.id
		dist := qn.dist

		visited[id] = dist
		cur := valves[id]

		for t := range cur.tunnels {
			if t == start || (visited[t] > 0 && visited[t] < dist+1) {
				continue
			}

			q.Push(QueueNode{t, dist + 1})
		}
	}

	newTunnels := map[string]int{}

	for key, val := range visited {
		if valves[key].flow > 0 && key != start {
			newTunnels[key] = val
		}
	}

	v := valves[start]
	v.tunnels = newTunnels

	return v
}

func Part1(input string) int {
	valves := parseInput(input)
	newValves := map[string]Valve{
		"AA": getVertexes(valves, "AA"),
	}

	for key, val := range valves {
		if val.flow > 0 {
			newValves[key] = getVertexes(valves, key)
		}
	}

	fmt.Println(newValves)

	best := bestPath(newValves, map[string]bool{}, newValves["AA"], 30, 0)

	fmt.Println(best)

	return best
}

func Part2(input string) int {
	valves := parseInput(input)

	fmt.Println(valves)

	best := bestPath(valves, map[string]bool{}, valves["AA"], 30, 0)

	fmt.Println(best)

	return best
}

func Run(input string) {
	fmt.Println("Day 16 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
