package day12

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Point struct {
	x int
	y int
}

func parseInput(input string) (grid [][]int, start, end Point) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid = make([][]int, len(lines))

	for y, line := range lines {
		grid[y] = make([]int, len(line))

		for x, char := range line {

			p := Point{x, y}

			if char == 'S' {
				start = p
				grid[y][x] = 0
			} else if char == 'E' {
				end = p
				grid[y][x] = 25
			} else {
				grid[y][x] = int(char - 'a')
			}
		}
	}

	return
}

func (p Point) String() string {
	return fmt.Sprintf("%v:%v", p.x, p.y)
}

func Neighbours(grid [][]int, node Point) (neighbours []Point) {
	h := len(grid)
	w := len(grid[0])

	// North
	if node.y > 0 {
		neighbours = append(neighbours, Point{node.x, node.y - 1})
	}
	// South
	if node.y < h-1 {
		neighbours = append(neighbours, Point{node.x, node.y + 1})
	}
	// West
	if node.x > 0 {
		neighbours = append(neighbours, Point{node.x - 1, node.y})
	}
	// East
	if node.x < w-1 {
		neighbours = append(neighbours, Point{node.x + 1, node.y})
	}

	return
}

type Node struct {
	point Point
	dist  int
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

// Dijkstra's
func fastestRoute(grid [][]int, start Point, check func(p Point) bool) int {
	distances := map[string]int{}
	visited := map[string]bool{}
	queue := &PriorityQueue{}
	heap.Init(queue)
	heap.Push(queue, Node{start, 0})

	for queue.Len() > 0 {
		cur := heap.Pop(queue).(Node)

		if visited[cur.point.String()] {
			continue
		}

		curVal := grid[cur.point.y][cur.point.x]
		curDist := cur.dist

		neighbours := Neighbours(grid, cur.point)

	neighboursLoop:
		for _, neighbour := range neighbours {
			nVal := grid[neighbour.y][neighbour.x]

			if curVal-nVal > 1 || visited[neighbour.String()] {
				continue neighboursLoop
			}
			nDist := distances[neighbour.String()]

			if nDist == 0 {
				nDist = math.MaxInt
			}

			if nDist > curDist+1 {
				nDist = curDist + 1
			}

			if check(neighbour) {
				return nDist
			}

			distances[neighbour.String()] = nDist
			heap.Push(queue, Node{neighbour, nDist})
		}

		visited[cur.point.String()] = true
	}

	return -1
}

func Part1(input string) int {
	grid, start, end := parseInput(input)
	dist := fastestRoute(grid, end, func(p Point) bool {
		return p.String() == start.String()
	})

	return dist
}

func Part2(input string) int {
	grid, _, end := parseInput(input)
	dist := fastestRoute(grid, end, func(p Point) bool {
		return grid[p.y][p.x] == 0
	})

	return dist

}

func Run(input string) {
	fmt.Println("Day 12 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
