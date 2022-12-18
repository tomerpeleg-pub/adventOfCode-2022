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

	// Allow 1 space outside the grid (needed for p2 floodfill)
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
		{point.x, point.y - 1, point.z}, //up
		{point.x, point.y + 1, point.z}, //down
		{point.x - 1, point.y, point.z}, //left
		{point.x + 1, point.y, point.z}, //right
		{point.x, point.y, point.z + 1}, //forward
		{point.x, point.y, point.z - 1}, //back
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

	// loop through every point on the grid
	for x := 0; x <= grid.w; x++ {
		for y := 0; y <= grid.h; y++ {
			for z := 0; z <= grid.d; z++ {
				p := Point{x, y, z}
				g := grid.Get(p)

				// if the point has a cube on it
				if g == FILLED {

					// loop through its' neighbours
					neighbours := grid.Neighbours(p)
					for _, n := range neighbours {

						// if the neighbour is not a cube
						// count it
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

	// Already visited, halt
	if myVal == val {
		return
	}

	// fill this points value
	grid.Set(point, val)

	// loop through all neighoburs
	neighbours := grid.Neighbours(point)
	for _, n := range neighbours {
		v := grid.Get(n)

		// if the neighbour is the same as this point (used to be)
		if v == myVal {
			// then continue filling from there
			grid.FloodFill(n, val)
		} else if v == FILLED {
			// count how many cubes we come across
			// (avoids having to loop over newly filled points afterwards)
			total++
		}
	}
}

func Part2(input string) int {
	grid := parseInput(input)
	total = 0 // Count of empty spaces

	// Starting at a corner outside the grid,
	// fill all the outside spaces with OUTSIDE=2
	// as you do, count any cubes that are found
	grid.FloodFill(Point{-1, -1, -1}, OUTSIDE)
	return total
}

func Run(input string) {
	fmt.Println("Day 18 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
