package day17

import (
	"fmt"
	"math/big"
	"strings"
)

const rocksInput = `
####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##`

const Wall Line = 0b100000001
const Floor Line = 0b111111111

type Line int

func (l Line) String() string {
	if l == Floor {
		return "+-------+"
	}

	str := ""
	lStr := fmt.Sprintf("%08b", l)

	for i, char := range lStr {
		if i == 0 || i == 8 {
			str += "|"
		} else if char == '1' {
			str += "#"
		} else {
			str += "."
		}
	}

	return str
}

type Rock []Line

func (r Rock) Shift(n Move) Rock {
	newRock := make(Rock, len(r))

	for i, l := range r {
		if n < 0 {
			newRock[i] = l << -n
		} else {
			newRock[i] = l >> n
		}
	}

	return newRock
}

func (r Rock) String() string {
	str := ""

	for _, line := range r {
		str += line.String() + "\n"
	}

	return str
}

type Grid []Line

func (g Grid) String() string {
	str := "\n"

	for i := len(g) - 1; i >= 0; i-- {
		line := g[i]
		str += fmt.Sprint(i) + ": " + line.String() + "\n"
	}

	return str
}

const (
	LEFT  = -1
	RIGHT = 1
)

type Move int

func (m Move) String() string {
	if m == 1 {
		return ">"
	} else {
		return "<"
	}
}

func parseRock(input string) Rock {
	rock := Rock{}

	for _, line := range strings.Split(input, "\n") {
		var l Line = 0

		j, char := 0, rune(0)
		le := len(line)

		for j, char = range line {
			if char == '#' {
				l |= 1 << (le - j)
			}
		}

		l = l << (6 - j)

		rock = append(rock, l)
	}

	return rock
}

func parseRocks(input string) []Rock {
	rocksStr := strings.Split(strings.TrimSpace(input), "\n\n")

	rocks := []Rock{}

	for _, rockStr := range rocksStr {
		rocks = append(rocks, parseRock(rockStr))
	}

	return rocks
}

func parseInput(input string) []Move {
	cleaned := strings.TrimSpace(input)
	moves := make([]Move, len(cleaned))

	for i, char := range cleaned {
		if char == '>' {
			moves[i] = 1
		} else {
			moves[i] = -1
		}
	}

	return moves
}

func testLine(lineA Line, lineB Line) bool {
	return lineA&lineB == 0
}

func testRock(rock Rock, grid Grid, y int) bool {
	for l, line := range rock {
		tl := Wall

		if y-l >= 0 && y-l < len(grid) {
			tl = grid[y-l]
		}

		if !testLine(line, tl) {
			return false
		}
	}

	return true
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func placeRock(rock Rock, grid Grid, y int) Grid {
	newGrid := make(Grid, Max(len(grid), y+1))
	copy(newGrid, grid)

	for l, line := range rock {
		if newGrid[y-l] == 0 {
			newGrid[y-l] = Wall
		}

		newGrid[y-l] |= line
	}
	return newGrid
}

func dropRock(grid Grid, rock Rock, moves []Move, m int, y int) (Grid, int, int) {
	y += len(rock) - 1
	// x := 2
	newRock := rock.Shift(2)
	// fmt.Println("Start:\n", placeRock(newRock, grid, y))

	for {
		test := newRock.Shift(moves[m%len(moves)])

		if testRock(test, grid, y) {
			newRock = test
		}
		// fmt.Println("After shift:", moves[m%len(moves)], "\n", placeRock(newRock, grid, y))

		m++

		if testRock(newRock, grid, y-1) {
			y--
			// fmt.Println("After move:\n", placeRock(newRock, grid, y))
		} else {
			return placeRock(newRock, grid, y), m, y
		}
	}
}

func Part1(input string) int {
	rocks := parseRocks(rocksInput)
	moves := parseInput(input)

	grid := Grid{Floor}
	m := 0

	for i := 0; i < 2022; i++ {
		grid, m, _ = dropRock(grid, rocks[i%len(rocks)], moves, m, len(grid)+3)
		m = m % len(moves)
	}

	return len(grid) - 1
}

func Part2(input string) int {
	rocks := parseRocks(rocksInput)
	moves := parseInput(input)

	grid := Grid{Floor}
	m := 0

	for i := 0; i < 3600; i++ {
		m = m % len(moves)
		r := i % len(rocks)
		grid, m, _ = dropRock(grid, rocks[r], moves, m, len(grid)+3)
	}

	// 1. Print out grid after a large number of shapes (e.g. 10000)
	// 2. Find a pattern in the grid, mine repeated every
	//    1730 shapes after shape 1006
	// 3. height for start of pattern = n * 2659 + 1006 (for me)
	// 4. Figure out what shape in the pattern the input is
	// 5. Run for that many iterations, then do maths below:

	a := big.NewInt(578034681)
	b := big.NewInt(2659)
	c := big.NewInt(1006)
	d := big.NewInt(int64(len(grid) - 3665))
	e := a.Mul(a, b)
	e = e.Add(e, c)
	e = e.Add(e, d)

	return int(e.Int64()) - 1
}

func Run(input string) {
	fmt.Println("Day 17 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
