package day8

import (
	"fmt"
	"strconv"
	"strings"
)

type Tree struct {
	Val     int
	Visible bool
	x       int
	y       int
}

func (t Tree) String() string {
	if t.Visible {
		return fmt.Sprintf("+%v+", t.Val)
	}
	return fmt.Sprintf("-%v-", t.Val)
}

func parseInput(input string) [][]Tree {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]Tree, len(lines))

	for yi, line := range lines {
		grid[yi] = make([]Tree, len(line))

		for xi, height := range line {
			val, _ := strconv.Atoi(string(height))
			grid[yi][xi] = Tree{
				Val: val,
				x:   xi,
				y:   yi,
			}
		}
	}

	return grid
}

func Part1(input string) int {
	grid := parseInput(input)

	count := 0

	height := len(grid)
	width := len(grid[0])

	hTop := make([]int, height)
	hBottom := make([]int, height)

	for i := range hTop {
		hTop[i] = -1
		hBottom[i] = -1
	}

	for y, row := range grid {

		hLeft := -1
		hRight := -1

		for x, val := range row {

			// left
			l := val
			if l.Val > hLeft {
				hLeft = l.Val
				if !l.Visible {
					count++
				}
				l.Visible = true
				grid[l.y][l.x] = l
			}

			// right
			r := row[width-x-1]
			if r.Val > hRight {
				hRight = r.Val
				if !r.Visible {
					count++
				}
				r.Visible = true
				grid[r.y][r.x] = r
			}

			// top
			t := grid[y][x]
			if t.Val > hTop[t.x] {
				hTop[t.x] = t.Val
				if !t.Visible {
					count++
				}
				t.Visible = true
				grid[t.y][t.x] = t
			}

			// bottom
			b := grid[height-y-1][x]
			if b.Val > hBottom[b.x] {
				hBottom[b.x] = b.Val
				if !b.Visible {
					count++
				}
				b.Visible = true
				grid[b.y][b.x] = b
			}
		}
	}

	return count
}

func Part2(input string) int {
	return 12
}

func Run(input string) {
	fmt.Println("Day 8 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
