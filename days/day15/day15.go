package day15

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("%v:%v", p.x, p.y)
}

func (a Point) Distance(b Point) int {
	distX := a.x - b.x
	distY := a.y - b.y

	if distX < 0 {
		distX = -distX
	}

	if distY < 0 {
		distY = -distY
	}

	return distX + distY
}

func (a Point) Equals(b Point) bool {
	return a.x == b.x && a.y == b.y
}

type Scanner struct {
	pos    Point
	beacon Point
	dist   int
}

func parseSensor(line string) Scanner {
	numsReg := regexp.MustCompile(`[-\d]+`)
	str := numsReg.FindAllString(line, -1)
	vals := [4]int{}

	for i, s := range str {
		vals[i], _ = strconv.Atoi(s)
	}

	scanner := Point{vals[0], vals[1]}
	beacon := Point{vals[2], vals[3]}
	dist := scanner.Distance(beacon)

	return Scanner{scanner, beacon, dist}
}

func parseInput(input string) []Scanner {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	scanners := []Scanner{}

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		scan := parseSensor(line)
		scanners = append(scanners, scan)
	}

	return scanners
}

func Part1(input string) int {
	scanners := parseInput(input)

	minX := math.MaxInt
	maxX := math.MinInt

	for _, scanner := range scanners {
		if scanner.pos.x-scanner.dist-1 < minX {
			minX = scanner.pos.x - scanner.dist - 1
		} else if scanner.pos.x+scanner.dist+1 > maxX {
			maxX = scanner.pos.x + scanner.dist - 1
		}
		if scanner.beacon.x < minX {
			minX = scanner.beacon.x
		} else if scanner.beacon.x > maxX {
			maxX = scanner.beacon.x
		}
	}

	fmt.Println(minX, maxX)

	x := 0
	count := 0

xLoop:
	for x = minX - 1; x <= maxX+1; x++ {

		for _, scanner := range scanners {

			p := Point{x, 2000000}
			dist := p.Distance(scanner.pos)

			if !p.Equals(scanner.beacon) && dist <= scanner.dist {
				count++
				continue xLoop
			}
		}
	}

	return count
}

func Min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func Part2(input string) int {
	scanners := parseInput(input)

	minX := 0
	maxX := 20

	// count := 0

	for i, sa := range scanners {

		// points
		left := sa.pos.x - sa.dist - 1
		right := sa.pos.x + sa.dist + 1
		top := sa.pos.y - sa.dist - 1
		bottom := sa.pos.y + sa.dist + 1

	yLoop:
		for y := 0; y <= sa.dist; y++ {

			// Looping around the boundary
			p1 := Point{left + y, sa.pos.y - y}
			p2 := Point{right - y, sa.pos.y - y}
			p3 := Point{sa.pos.x + y, top + y}
			p4 := Point{sa.pos.x + y, bottom - y}

			if p1.x < minX || p1.x > maxX || p1.y < minX || p1.y > maxX {
				continue yLoop
			}

		sLoop:
			for j, sb := range scanners {
				if i == j {
					continue sLoop
				}

				dist1 := p1.Distance(sb.pos)
				dist2 := p2.Distance(sb.pos)
				dist3 := p3.Distance(sb.pos)
				dist4 := p4.Distance(sb.pos)

				if dist1 <= sb.dist {
					continue yLoop
				} else if dist2 <= sb.dist {
					continue yLoop
				} else if dist3 <= sb.dist {
					continue yLoop
				} else if dist4 <= sb.dist {
					continue yLoop
				}
			}

			fmt.Println("Found!", p1, p2, p3, p4)
			return (p1.x * 4000000) + p1.y
		}

	}
	return -1

	// for y := minX; y <= maxX; y++ {
	// xLoop:
	// 	for x := minX; x <= maxX; x++ {
	// 		p := Point{x, y}

	// 		for _, scanner := range scanners {

	// 			dist := p.Distance(scanner.pos)

	// 			if dist <= scanner.dist {
	// 				continue xLoop
	// 			}
	// 		}

	// 		return (x * 4000000) + y
	// 	}
	// }

	// return count
}

// 4396473
func Run(input string) {
	fmt.Println("Day 15 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
