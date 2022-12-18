package day18

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Point struct {
	x, y, z int
}

func (p Point) String() string {
	return fmt.Sprintf("{%v:%v:%v}", p.x, p.y, p.z)
}

type Grid struct {
	points map[string]int

	w int
	h int
	d int
}

const (
	OUT     = -1
	EMPTY   = 0
	FILLED  = 1
	OUTSIDE = 2
)

func (grid Grid) Get(point Point) int {
	if point.x < -1 || point.y < -1 || point.z < -1 ||
		point.x > 1+grid.w || point.y > 1+grid.h || point.z > 1+grid.d {
		return OUT
	}

	return grid.points[point.String()]
}

func (grid *Grid) Set(point Point, val int) {
	grid.points[point.String()] = val
}

func (grid Grid) Neighbours(point Point) []Point {
	return []Point{
		//up
		{point.x, point.y - 1, point.z},
		//down
		{point.x, point.y + 1, point.z},
		//left
		{point.x - 1, point.y, point.z},
		//right
		{point.x + 1, point.y, point.z},
		//forward
		{point.x, point.y, point.z + 1},
		//back
		{point.x, point.y, point.z - 1},
	}
}

func parseInput(input string) Grid {

	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	maxX := 0
	maxY := 0
	maxZ := 0

	points := map[string]int{}

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])

		p := Point{x, y, z}
		points[p.String()] = FILLED

		maxX = Max(maxX, x)
		maxY = Max(maxY, y)
		maxZ = Max(maxZ, z)
	}

	return Grid{points, maxX, maxY, maxZ}
}

func Part1(input string) int {
	grid := parseInput(input)

	count := 0

	for x := 0; x <= grid.w; x++ {
		for y := 0; y <= grid.h; y++ {
			for z := 0; z <= grid.d; z++ {
				p := Point{x, y, z}
				g := grid.Get(p)

				if g == FILLED {
					neighbours := grid.Neighbours(p)

					for _, n := range neighbours {
						if grid.Get(n) != FILLED {
							count++
						}
					}
				}
			}
		}
	}

	return count
}

var total int = 0

func (grid *Grid) FloodFill(point Point, val int) {
	myVal := grid.Get(point)
	if myVal == val {
		return
	}

	neighbours := grid.Neighbours(point)
	grid.Set(point, val)

	for _, n := range neighbours {
		v := grid.Get(n)

		if v == myVal {
			grid.FloodFill(n, val)
		} else if v == FILLED {
			total++
		}
	}
}

func Part2(input string) int {
	total = 0
	grid := parseInput(input)
	grid.FloodFill(Point{0, 0, 0}, OUTSIDE)
	return total
}

func Run(input string) {
	fmt.Println("Day 18 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
