package day24

import (
	"bufio"
	"container/heap"
	"fmt"
	"strings"
)

const (
	N    = 0 - 1i
	E    = 1 + 0i
	S    = 0 + 1i
	W    = -1 + 0i
	NONE = 0 + 0i
)

const (
	WALL     = -1
	EMPTY    = 0
	BLIZZARD = 1
)

type Blizzard struct {
	pos complex128
	dir complex128
}

type Grid map[complex128]int

type State struct {
	grid      Grid
	blizzards []Blizzard
	w         float64
	h         float64
	start     complex128
	end       complex128
}

func (s State) String() string {
	fmt.Println("Printing grid:", s.w, s.h)
	str := ""

	for y := 0.0; y < s.h; y++ {
		for x := 0.0; x < s.w; x++ {
			p := complex(x, y)
			switch s.grid[p] {
			case WALL:
				str += "#"
			case EMPTY:
				str += "."
			case BLIZZARD:
				str += "o"
			default:
				str += fmt.Sprint(s.grid[p])
			}
		}
		str += "\n"
	}

	return str
}

func parseInput(input string) State {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	grid := Grid{}
	blizzards := []Blizzard{}

	y := 0

	lastEmpty := 0 + 0i

	for scanner.Scan() {
		line := scanner.Text()

		for x, char := range line {
			p := complex(float64(x), float64(y))

			switch char {
			case '#':
				grid[p] = WALL
			case '.':
				grid[p] = EMPTY
				lastEmpty = p
			case '^':
				grid[p] = BLIZZARD
				blizzards = append(blizzards, Blizzard{p, N})
			case '>':
				grid[p] = BLIZZARD
				blizzards = append(blizzards, Blizzard{p, E})
			case 'v':
				grid[p] = BLIZZARD
				blizzards = append(blizzards, Blizzard{p, S})
			case '<':
				grid[p] = BLIZZARD
				blizzards = append(blizzards, Blizzard{p, W})
			}
		}

		y++
	}

	start := complex(1, 0)
	end := lastEmpty

	grid[start+N] = WALL
	grid[end+S] = WALL

	return State{grid, blizzards, real(end) + 2, imag(end) + 1, start, end}
}

func simulate(state State) State {
	g2 := Grid{}
	b2 := make([]Blizzard, len(state.blizzards))

	for pos, cell := range state.grid {
		g2[pos] = cell
	}

	for i, blizzard := range state.blizzards {
		g2[blizzard.pos]--
		nb := Blizzard{
			pos: blizzard.pos + blizzard.dir,
			dir: blizzard.dir,
		}
		if g2[nb.pos] == WALL {
			switch nb.dir {
			case N:
				nb.pos += complex(0, state.h-2)
			case S:
				nb.pos -= complex(0, state.h-2)
			case E:
				nb.pos -= complex(state.w-2, 0)
			case W:
				nb.pos += complex(state.w-2, 0)
			}
		}
		g2[nb.pos]++
		b2[i] = nb
	}

	return State{g2, b2, state.w, state.h, state.start, state.end}
}

type Node struct {
	dist int
	pos  complex128
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}
func (pq *PriorityQueue) Pop() (popped interface{}) {
	popped = (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return
}

func findPath(initial State) (int, State) {

	queue := &PriorityQueue{}
	heap.Init(queue)
	heap.Push(queue, Node{dist: 0, pos: initial.start})

	cube := []State{initial}

	visited := []map[complex128]bool{}

	for queue.Len() > 0 {
		node := heap.Pop(queue).(Node)
		if node.dist >= len(visited) {
			visited = append(visited, map[complex128]bool{})
		}
		if visited[node.dist][node.pos] {
			continue
		}
		if node.pos == initial.end {
			return node.dist, cube[node.dist]
		}

		visited[node.dist][node.pos] = true

		if node.dist+1 >= len(cube) {
			cube = append(cube, simulate(cube[len(cube)-1]))
		}

		layer := cube[node.dist+1]
		dirs := [5]complex128{N, E, S, W, NONE}

		for _, dir := range dirs {
			np := node.pos + dir

			if layer.grid[np] == EMPTY {
				heap.Push(queue, Node{pos: np, dist: node.dist + 1})
			}
		}
	}

	return -1, initial
}

func Part1(input string) int {
	state := parseInput(input)

	time, _ := findPath(state)
	return time
}

func Part2(input string) int {
	state := parseInput(input)

	// to the end
	s1, s1State := findPath(state)
	s1State.start, s1State.end = s1State.end, s1State.start

	// end to start
	s2, s2State := findPath(s1State)
	s2State.start, s2State.end = s2State.end, s2State.start

	// start to end again
	s3, _ := findPath(s2State)

	return s1 + s2 + s3
}

func Run(input string) {
	fmt.Println("Day 24 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
