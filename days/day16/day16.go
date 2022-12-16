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

func bestPath(valves map[string]Valve, opened map[string]bool, start Valve, mins int, total int) (int, string) {
	if mins <= 1 {
		return total, ""
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
	bestNode := ""

	for id, dist := range start.tunnels {
		if dist > time {
			continue
		}
		if newOpened[id] {
			continue
		}
		v := valves[id]

		result, _ := bestPath(valves, newOpened, v, time-dist, curTotal)

		if result > curHighest {
			curHighest = result
			bestNode = id
		}
	}

	return curHighest, bestNode
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

func getPairs(input string) map[string]Valve {
	valves := parseInput(input)
	newValves := map[string]Valve{
		"AA": getVertexes(valves, "AA"),
	}

	for key, val := range valves {
		if val.flow > 0 {
			newValves[key] = getVertexes(valves, key)
		}
	}

	return newValves
}

func Part1(input string) int {
	valves := getPairs(input)
	best, node := bestPath(valves, map[string]bool{}, valves["AA"], 30, 0)

	fmt.Println(best, "node:", node)

	return best
}

type Player struct {
	node Valve
	wait int
}

func calcScore(valves map[string]Valve, opened map[string]bool) int {
	total := 0

	op := []string{}

	for id, opened := range opened {
		if opened {
			total += valves[id].flow
			op = append(op, id)
		}
	}

	fmt.Println("Open valves:", op, "pressure:", total)

	return total
}

// Not working
func bestTwoPlayerPath(valves map[string]Valve, start Valve, mins int) int {
	p1 := Player{start, 0}
	p2 := Player{start, 0}
	opened := map[string]bool{}
	curScore := 0

	for i := 1; i < mins; i++ {
		fmt.Println("== Minute", i, "==")
		newOpened := copyMap(opened)
		newOpened[p1.node.id] = true
		newOpened[p2.node.id] = true

		// p1
		if p1.wait == 1 {
			opened[p1.node.id] = true
			fmt.Println("Min", i, "Player 1 opens valve", p1.node.id)
			p1.wait--
		} else if p1.wait == 0 {
			_, node := bestPath(valves, newOpened, p1.node, mins-i, 0)

			if node != p1.node.id && node != "" {
				p1.wait = valves[p1.node.id].tunnels[node]
				fmt.Println("Min", i, "Player 1 moves from", p1.node.id, "to", node, ", taking", p1.wait, "days")
				p1.node = valves[node]
				newOpened[node] = true
			} else {
				p1.wait = 500
			}
		} else {
			fmt.Println("Min", i, "Player 1 on way to", p1.node.id)
			p1.wait--
		}

		// p2
		if p2.wait == 1 {
			opened[p2.node.id] = true
			fmt.Println("Min", i, "Player 2 opens valve", p2.node.id)
			p2.wait--
		} else if p2.wait == 0 {
			_, node := bestPath(valves, newOpened, p2.node, mins-i, 0)

			if node != p2.node.id && node != "" {
				p2.wait = valves[p2.node.id].tunnels[node]
				fmt.Println("Min", i, "Player 2 moves from", p2.node.id, "to", node, ", taking", p2.wait, "days")
				p2.node = valves[node]
			} else {
				p2.wait = 500
			}
		} else {
			fmt.Println("Min", i, "Player 2 on way to", p2.node.id)
			p2.wait--
		}

		curScore += calcScore(valves, opened)
	}

	return curScore
}

func Part2(input string) int {
	valves := getPairs(input)
	fmt.Println(valves)
	best := bestTwoPlayerPath(valves, valves["AA"], 26)

	fmt.Println(best)

	return best
}

func Run(input string) {
	fmt.Println("Day 16 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
