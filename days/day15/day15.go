package day15

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Rect struct{ x, y, mx, my int }

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func dist(a complex128, b complex128) float64 {
	return abs(real(a)-real(b)) + abs(imag(a)-imag(b))
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInput(input string) (Rect, map[complex128]bool, map[complex128]float64) {
	sensorRegex := regexp.MustCompile(`Sensor at x=([-\d]+), y=([-\d]+): closest beacon is at x=([-\d]+), y=([-\d]+)`)
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	beacons := map[complex128]bool{}
	sensors := map[complex128]float64{}

	rect := Rect{math.MaxInt, math.MaxInt, math.MinInt, math.MinInt}

	for scanner.Scan() {
		line := scanner.Text()
		match := sensorRegex.FindAllStringSubmatch(line, -1)
		sx, sy, bx, by := match[0][1], match[0][2], match[0][3], match[0][4]
		sxi, _ := strconv.Atoi(sx)
		syi, _ := strconv.Atoi(sy)
		bxi, _ := strconv.Atoi(bx)
		byi, _ := strconv.Atoi(by)

		sensor := complex(float64(sxi), float64(syi))
		beacon := complex(float64(bxi), float64(byi))

		beacons[beacon] = true
		sensors[sensor] = dist(sensor, beacon)

		rect.x = min(rect.x, sxi-int(sensors[sensor]))
		rect.mx = max(rect.mx, sxi+int(sensors[sensor]))
		rect.y = min(rect.y, syi-int(sensors[sensor]))
		rect.my = max(rect.my, syi+int(sensors[sensor]))
	}

	return rect, beacons, sensors
}

func scanned(sensors map[complex128]float64, p complex128) bool {
	for pos, r := range sensors {
		d := dist(p, pos)

		if d <= r {
			return true
		}
	}

	return false
}

func Part1(input string) int {
	rect, beacons, sensors := parseInput(input)
	tot := 0
	y := 10.0

	fmt.Println(rect)

	for x := rect.x; x <= rect.mx; x++ {
		p := complex(float64(x), y)

		if beacons[p] {
			continue
		}

		if scanned(sensors, p) {
			tot++
		}
	}

	return tot
}

func inBounds(a complex128) bool {
	return real(a) >= 0 && real(a) <= 4000000 && imag(a) >= 0 && imag(a) <= 4000000
}

func Part2(input string) int {
	_, beacons, sensors := parseInput(input)

	dirs := [4]complex128{1 - 1i, 1 + 1i, -1 - 1i, -1 + 1i}

	for pos, r := range sensors {

		// nw
		nw := complex(real(pos)-r-1, imag(pos))
	nwLoop:
		for dist(nw, pos) == r+1 {
			if !inBounds(nw) {

				nw += dirs[0]
				continue nwLoop
			}
			if !beacons[nw] && !scanned(sensors, nw) {
				fmt.Println("Found! nw", nw)
				return int(real(nw)*4000000 + imag(nw))
			}
			nw += dirs[0]
		}
		// ne
		ne := complex(real(pos), imag(pos)-r-1)
	neLoop:
		for dist(ne, pos) == r+1 {

			if !inBounds(ne) {

				ne += dirs[1]
				continue neLoop
			}
			if !beacons[ne] && !scanned(sensors, ne) {
				fmt.Println("Found! ne", ne)
				return int(real(ne)*4000000 + imag(ne))
			}
			ne += dirs[1]
		}
		// se
		se := complex(real(pos)+r+1, imag(pos))
	seLoop:
		for dist(se, pos) == r+1 {

			if !inBounds(se) {

				se += dirs[2]
				continue seLoop
			}
			if !beacons[se] && !scanned(sensors, se) {
				fmt.Println("Found! se", se)
				return int(real(se)*4000000 + imag(se))
			}
			se += dirs[2]
		}
		// sw
		sw := complex(real(pos), imag(pos)+r+1)
	swLoop:
		for dist(sw, pos) == r+1 {

			if !inBounds(sw) {

				sw += dirs[3]
				continue swLoop
			}
			if !beacons[sw] && !scanned(sensors, sw) {
				fmt.Println("Found! sw", sw)
				return int(real(sw)*4000000 + imag(sw))
			}
			sw += dirs[3]
		}

		fmt.Println("Ruled out sensor at", pos)
	}

	return -1
}

func Run(input string) {
	fmt.Println("Day 15 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
