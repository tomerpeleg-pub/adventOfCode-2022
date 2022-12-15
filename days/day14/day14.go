package day14

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	EMPTY = 0
	WALL  = 1
	SAND  = 2
)

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("%v:%v", p.x, p.y)
}

type Boundary struct {
	left   int
	right  int
	bottom int
	top    int
}

func (b Boundary) String() string {
	return fmt.Sprintf("{ l: %v, r: %v, b: %v, t: %v }", b.left, b.right, b.bottom, b.top)
}

func parseCoord(coord string) Point {
	pStrings := strings.Split(coord, ",")
	x, _ := strconv.Atoi(pStrings[0])
	y, _ := strconv.Atoi(pStrings[1])

	return Point{x, y}
}

func (boundary Boundary) Compare(point Point) Boundary {
	if point.x < boundary.left {
		boundary.left = point.x
	} else if point.x > boundary.right {
		boundary.right = point.x
	}

	if point.y < boundary.top {
		boundary.top = point.y
	} else if point.y > boundary.bottom {
		boundary.bottom = point.y
	}
	return boundary
}

type Grid map[string]int

func (grid Grid) String(boundary Boundary) string {
	str := ""

	for y := boundary.top; y <= boundary.bottom; y++ {
		for x := boundary.left; x <= boundary.right; x++ {
			point := Point{x, y}

			if grid[point.String()] == WALL {
				str += "#"
			} else if grid[point.String()] == SAND {
				str += "o"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func parseInput(input string) (points Grid, boundary Boundary) {
	points = Grid{}
	boundary = Boundary{
		left:   math.MaxInt,
		right:  math.MinInt,
		bottom: math.MinInt,
		top:    math.MaxInt,
	}

	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, " -> ")

		prev := parseCoord(coords[0])
		boundary = boundary.Compare(prev)

		for _, coord := range coords[1:] {
			point := parseCoord(coord)
			boundary = boundary.Compare(point)

			if point.x != prev.x {
				if point.x > prev.x {
					for j := prev.x; j <= point.x; j++ {
						points[Point{x: j, y: point.y}.String()] = WALL
					}
				} else {
					for j := point.x; j <= prev.x; j++ {
						points[Point{x: j, y: point.y}.String()] = WALL
					}
				}
			} else if point.y != prev.y {
				if point.y > prev.y {
					for j := prev.y; j <= point.y; j++ {
						points[Point{x: point.x, y: j}.String()] = WALL
					}
				} else {
					for j := point.y; j <= prev.y; j++ {
						points[Point{x: point.x, y: j}.String()] = WALL
					}
				}
			}
			prev = point
		}
	}

	return
}

func DropSand(points Grid, boundary Boundary) (Grid, Boundary, bool) {
	sand := Point{500, 0}
	next := sand

	if points[sand.String()] != EMPTY {
		return points, boundary, false
	}

	for {
		if next.x < boundary.left || next.x > boundary.right || next.y > boundary.bottom {
			return points, boundary, false
		}

		next = Point{sand.x, sand.y + 1}
		if points[next.String()] == EMPTY {
			sand = next
			continue
		}

		next = Point{sand.x - 1, sand.y + 1}
		if points[next.String()] == EMPTY {
			sand = next
			continue
		}

		next = Point{sand.x + 1, sand.y + 1}
		if points[next.String()] == EMPTY {
			sand = next
			continue
		}

		points[sand.String()] = SAND
		boundary = boundary.Compare(sand)
		return points, boundary, true
	}
}

func Part1(input string) int {
	points, boundary := parseInput(input)
	cont := true

	fmt.Println("Before:")
	fmt.Println(boundary)
	fmt.Println(points.String(boundary))

	for i := 0; cont; i++ {
		points, boundary, cont = DropSand(points, boundary)

		if !cont {
			fmt.Println("After:")
			fmt.Println(points.String(boundary))
			return i
		}
	}

	return -1
}

func Part2(input string) int {
	points, boundary := parseInput(input)
	boundary.left -= 150  // play around with this num
	boundary.right += 150 // play around with this num
	boundary.bottom += 2
	cont := true

	for x := boundary.left; x <= boundary.right; x++ {
		points[Point{x, boundary.bottom}.String()] = WALL
	}

	fmt.Println("Before:")
	fmt.Println(boundary)

	for i := 0; cont; i++ {
		points, boundary, cont = DropSand(points, boundary)

		if !cont {
			fmt.Println("After:", i, boundary)
			// uncomment this if you want to cry
			// fmt.Println(points.String(boundary))
			return i
		}
	}

	return -1
}

func Run(input string) {
	fmt.Println("Day 14 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
