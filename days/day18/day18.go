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

type Grid struct {
	cubes [][][]bool
	w     int
	h     int
	d     int
}

func parseInput(input string) ([]Point, Grid) {

	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	maxX := 0
	maxY := 0
	maxZ := 0

	points := []Point{}

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")

		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])

		p := Point{x, y, z}

		points = append(points, p)

		maxX = Max(maxX, x+1)
		maxY = Max(maxY, y+1)
		maxZ = Max(maxZ, z+1)
	}

	grid := make([][][]bool, maxX)

	for x := 0; x < maxX; x++ {
		grid[x] = make([][]bool, maxY)

		for y := 0; y < maxY; y++ {
			grid[x][y] = make([]bool, maxZ)
		}
	}

	for _, p := range points {
		grid[p.x][p.y][p.z] = true
	}

	return points, Grid{
		grid, maxX, maxY, maxZ,
	}
}

func countSides(grid Grid, p Point) int {
	count := 0

	if p.x == 0 || !grid.cubes[p.x-1][p.y][p.z] {
		count++
	}
	if p.x == grid.w-1 || !grid.cubes[p.x+1][p.y][p.z] {
		count++
	}

	if p.y == 0 || !grid.cubes[p.x][p.y-1][p.z] {
		count++
	}
	if p.y == grid.h-1 || !grid.cubes[p.x][p.y+1][p.z] {
		count++
	}

	if p.z == 0 || !grid.cubes[p.x][p.y][p.z-1] {
		count++
	}
	if p.z == grid.d-1 || !grid.cubes[p.x][p.y][p.z+1] {
		count++
	}

	return count
}

func countOutsideEdges(grid Grid, p Point) int {
	count := 0

	if p.x == 0 {
		count++
	}
	if p.x == grid.w-1 {
		count++
	}

	if p.y == 0 {
		count++
	}
	if p.y == grid.h-1 {
		count++
	}

	if p.z == 0 {
		count++
	}
	if p.z == grid.d-1 {
		count++
	}

	return count
}

func Part1(input string) int {
	points, grid := parseInput(input)
	total := 0

	for _, point := range points {
		total += countSides(grid, point)
	}

	return total
}

func findNeighbours(grid Grid, p Point) []Point {
	neighbours := []Point{}

	if p.x > 0 && !grid.cubes[p.x-1][p.y][p.z] {
		neighbours = append(neighbours, Point{p.x - 1, p.y, p.z})
	}
	if p.x < grid.w-1 && !grid.cubes[p.x+1][p.y][p.z] {
		neighbours = append(neighbours, Point{p.x + 1, p.y, p.z})
	}

	if p.y > 0 && !grid.cubes[p.x][p.y-1][p.z] {
		neighbours = append(neighbours, Point{p.x, p.y - 1, p.z})
	}
	if p.y < grid.h-1 && !grid.cubes[p.x][p.y+1][p.z] {
		neighbours = append(neighbours, Point{p.x, p.y + 1, p.z})
	}

	if p.z > 0 && !grid.cubes[p.x][p.y][p.z-1] {
		neighbours = append(neighbours, Point{p.x, p.y, p.z - 1})
	}
	if p.z < grid.d-1 && !grid.cubes[p.x][p.y][p.z+1] {
		neighbours = append(neighbours, Point{p.x, p.y, p.z + 1})
	}

	return neighbours
}

func floodFill(grid Grid, p Point) {
	if grid.cubes[p.x][p.y][p.z] {
		return
	}

	grid.cubes[p.x][p.y][p.z] = true
	emptyNeighbours := findNeighbours(grid, p)

	for _, en := range emptyNeighbours {
		floodFill(grid, en)
	}
}

func counFilledtSides(grid Grid, p Point) int {
	count := 0

	if p.x == 0 || grid.cubes[p.x-1][p.y][p.z] {
		count++
	}
	if p.x == grid.w-1 || grid.cubes[p.x+1][p.y][p.z] {
		count++
	}

	if p.y == 0 || grid.cubes[p.x][p.y-1][p.z] {
		count++
	}
	if p.y == grid.h-1 || grid.cubes[p.x][p.y+1][p.z] {
		count++
	}

	if p.z == 0 || grid.cubes[p.x][p.y][p.z-1] {
		count++
	}
	if p.z == grid.d-1 || grid.cubes[p.x][p.y][p.z+1] {
		count++
	}

	return count
}

func Part2(input string) int {
	points, grid := parseInput(input)
	total := 0
	insideEdges := 0

	for _, point := range points {
		emptyNeighbours := findNeighbours(grid, point)

		for _, en := range emptyNeighbours {
			if countOutsideEdges(grid, en) > 0 {
				floodFill(grid, en)
			}
		}
	}

	for x := range grid.cubes {
		for y := range grid.cubes[x] {
			for z, filled := range grid.cubes[x][y] {
				p := Point{x, y, z}

				if !filled {
					insideEdges += counFilledtSides(grid, p)
				} else {
					total += countSides(grid, p)
				}
			}
		}
	}

	return Part1(input) - insideEdges
}

func Run(input string) {
	fmt.Println("Day 18 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
