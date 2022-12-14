package day12

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("%v:%v", p.x, p.y)
}

type Node struct {
	height     int
	point      Point
	dist       int
	neighbours []Node
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

func (node *Node) Neighbours(grid Grid) []Node {
	neighbours := []Node{}

	h := len(grid)
	w := len(grid[0])

	// North
	if node.point.y > 0 {
		north := grid[node.point.y-1][node.point.x]
		diff := north.height - node.height

		if diff <= 1 {
			neighbours = append(neighbours, north)
		}
	}

	// South
	if node.point.y < h-1 {
		south := grid[node.point.y+1][node.point.x]
		diff := south.height - node.height

		if diff <= 1 {
			neighbours = append(neighbours, south)
		}
	}

	// East
	if node.point.x > 0 {
		east := grid[node.point.y][node.point.x-1]
		diff := east.height - node.height

		if diff <= 1 {
			neighbours = append(neighbours, east)
		}
	}

	// West
	if node.point.x < w-1 {
		west := grid[node.point.y][node.point.x+1]
		diff := west.height - node.height

		if diff <= 1 {
			neighbours = append(neighbours, west)
		}
	}

	node.neighbours = neighbours

	return neighbours
}

type Grid [][]Node

func (grid Grid) String() string {
	str := ""

	for _, row := range grid {
		for _, node := range row {
			str += string(rune(node.height + 'a'))
		}
		str += "\n"
	}

	return str
}

func parseInput(input string) (grid Grid, start *Node, end *Node) {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []Node{})

		for x, char := range line {
			node := Node{point: Point{x, y}}

			if char == 'S' {
				node.height = 0
				node.dist = 0
				start = &node
			} else if char == 'E' {
				node.height = int('z' - 'a')
				node.dist = math.MaxInt32
				end = &node
			} else {
				node.dist = math.MaxInt32
				node.height = int(char - 'a')
			}

			grid[y] = append(grid[y], node)
		}

		y++
	}

	return grid, start, end
}

func fastestRoute(start Node, end Node, grid Grid) int {
	visited := map[string]bool{}
	queue := &PriorityQueue{}
	heap.Init(queue)

	for y, row := range grid {
		for x, col := range row {
			if col.height == 0 {
				grid[y][x].dist = 0
				queue.Push(grid[y][x])
			} else {
				grid[y][x].dist = math.MaxInt
			}
		}
	}

	for queue.Len() > 0 {
		// fmt.Println("Step:", queue)
		cur := heap.Pop(queue).(Node)
		if visited[cur.point.String()] {
			continue
		}

		dist := cur.dist

		if cur.neighbours == nil {
			cur.Neighbours(grid)
		}

	neighbourLoop:
		for _, neighbour := range cur.neighbours {
			if visited[neighbour.point.String()] {
				continue neighbourLoop
			}

			tDist := dist + 1

			if tDist < neighbour.dist {
				neighbour.dist = tDist
			}

			if neighbour.point.String() == end.point.String() {
				return tDist
			}

			heap.Push(queue, neighbour)
		}

		visited[cur.point.String()] = true
	}

	return math.MaxInt32
}

func Part1(input string) int {
	grid, start, end := parseInput(input)
	dist := fastestRoute(*start, *end, grid)

	return dist
}

func Part2(input string) int {
	grid, start, end := parseInput(input)

	lowest := fastestRoute(*start, *end, grid)
	count := 0

	// for _, row := range grid {
	// 	for _, node := range row {
	// 		if node.height == 0 {
	// 			dist
	// 			count++

	// 			if dist < lowest {
	// 				lowest = dist
	// 				fmt.Println("New lowest:", node.point, "dist", dist)
	// 			}
	// 		}
	// 	}
	// }
	fmt.Println("Checked:", count, "lowest:", lowest)

	return lowest

}

func Run(input string) {
	fmt.Println("Day 12 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
