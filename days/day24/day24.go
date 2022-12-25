package day24

import (
	"bufio"
	"container/heap"
	"fmt"
	"strings"
)

const (
	N = 0 + -1i
	E = 1 + 0i
	S = 0 + 1i
	W = -1 + 0i
)

const (
	WALL  = -1
	EMPTY = 0
	WIND  = 1
)

type Grid map[complex128]int
type Blizzard [2]complex128
type State struct {
	grid      Grid
	blizzards []Blizzard
	w         int
	h         int
	start     complex128
	end       complex128
	time      int
}

func (b Blizzard) String() string {
	switch b[1] {
	case N:
		return "^"
	case S:
		return "V"
	case E:
		return ">"
	case W:
		return "<"
	}
	return ""
}

func (s State) String() string {
	g := make([][]string, s.h)

	for _, b := range s.blizzards {
		x := int(real(b[0]))
		y := int(imag(b[0]))

		if g[y] == nil {
			g[y] = make([]string, s.w)
		}

		if s.grid[b[0]] == 1 {
			g[y][x] = b.String()
		} else {
			g[y][x] = fmt.Sprint(s.grid[b[0]])
		}
	}

	// fmt.Println("end", s.end)
	str := ""
	for y := 0; y < s.h; y++ {
		for x := 0; x < s.w; x++ {
			p := complex(float64(x), float64(y))

			if p == s.start {
				str += "S"
			} else if p == s.end {
				str += "E"
			} else if s.grid[p] == WALL {
				str += "#"
			} else {
				if g[y] != nil && g[y][x] != "" {
					str += g[y][x]
				} else {
					str += "."
				}
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

	y, w := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		w = len(line)

		for x, char := range line {
			p := complex(float64(x), float64(y))

			switch char {
			case '.':
				grid[p] = EMPTY
			case '#':
				grid[p] = WALL
			case '>':
				blizzards = append(blizzards, Blizzard{p, E})
				grid[p]++
			case '<':
				blizzards = append(blizzards, Blizzard{p, W})
				grid[p]++
			case 'v':
				blizzards = append(blizzards, Blizzard{p, S})
				grid[p]++
			case '^':
				blizzards = append(blizzards, Blizzard{p, N})
				grid[p]++
			}
		}

		y++
	}

	start := 1 + 0i
	end := complex(float64(w-2), float64(y-1))
	// fmt.Println(start, end)
	grid[start-1i] = WALL
	grid[end+1i] = WALL

	return State{
		grid, blizzards, w, y, start, end, 0,
	}
}

func simulate(s State) State {
	s2 := State{
		grid:      map[complex128]int{},
		blizzards: make([]Blizzard, len(s.blizzards)),
		w:         s.w,
		h:         s.h,
		time:      s.time + 1,
		start:     s.start,
		end:       s.end,
	}

	for i, w := range s.grid {
		if w == WALL {
			s2.grid[i] = w
		}
	}

	for i, b := range s.blizzards {
		b[0] += b[1]

		if isOut(s, b[0]) {
			switch b[1] {
			case N:
				b[0] += complex(0, float64(s2.h-2))
			case S:
				b[0] -= complex(0, float64(s2.h-2))
			case E:
				b[0] -= complex(float64(s2.w-2), 0)
			case W:
				b[0] += complex(float64(s2.w-2), 0)
			}
		}
		s2.blizzards[i] = b
		s2.grid[b[0]]++
	}

	return s2
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}

	return a
}

func dist(a complex128, b complex128) float64 {
	return abs(real(a)-real(b)) + abs(imag(a)-imag(b))
}

func isOut(s State, p complex128) bool {
	if p == s.start {
		return false
	}
	if real(p) <= 0 || imag(p) <= 0 || real(p) >= float64(s.w) || imag(p) >= float64(s.h) {
		return true
	}
	return false
}

type Node struct {
	pos   complex128
	dist  float64
	time  int
	moves []complex128
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// return pq[i].time*int(pq[i].dist) < pq[j].time*int(pq[j].dist)
	if pq[i].time == pq[j].time {
		return pq[i].dist < pq[j].dist

	} else {

		return pq[i].time < pq[j].time
	}
	// if pq[i].dist != pq[j].dist {
	// } else {
	// 	return pq[i].time < pq[j].time
	// }
	// return  <
	// }
	// return
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

func printMove(m complex128) string {
	switch m {
	case N:
		return "N"
	case S:
		return "S"
	case E:
		return "E"
	case W:
		return "W"
	}
	return ""
}

func (n Node) Norm() string {
	return fmt.Sprintf("{%v,%v}", n.pos, n.time)
}

func findPath(start State) int {
	fmt.Println("Finding path", start)

	start.time = 1
	sims := [500]State{start}

	for i := 1; i < 500; i++ {
		sims[i] = simulate(sims[i-1])
	}

	queue := make(PriorityQueue, 0)
	startNode := Node{
		pos:   start.start,
		dist:  dist(start.start, start.end),
		time:  1,
		moves: []complex128{},
	}
	heap.Push(&queue, startNode)

	visited := map[string]int{}

	// finish := Node{
	// 	time: math.MaxInt,
	// }

	for queue.Len() > 0 {
		cur := heap.Pop(&queue).(Node)
		if visited[cur.Norm()] > 0 {
			continue
		}
		visited[cur.Norm()] = cur.time
		sim := sims[cur.time]
		// fmt.Printf("== Checking day %v ===\n", cur.time)
		// fmt.Println(sim)

		dirs := [5]complex128{
			cur.pos,
			cur.pos + N,
			cur.pos + W,
			cur.pos + E,
			cur.pos + S,
		}

		for _, dir := range dirs {
			if dir == sim.end {
				fmt.Println("Found end, moves:", cur.time, cur.moves)

				return sim.time
			}

			// isStart := dir == 1+0i
			newState := Node{
				pos:   dir,
				dist:  dist(dir, sim.end),
				time:  sim.time,
				moves: append(cur.moves, dir),
			}

			if sim.grid[dir] == EMPTY {
				// fmt.Println("added one", newState.Norm())
				heap.Push(&queue, newState)
			}
		}
	}

	return -1
}

func Part1(input string) int {
	fmt.Println("== INITIAL ====")
	initialState := parseInput(input)
	fmt.Println(initialState)

	result := findPath(initialState)
	return result
}

func Part2(input string) int {
	return 12
}

func Run(input string) {
	fmt.Println("Day 24 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
