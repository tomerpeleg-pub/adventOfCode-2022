package day8

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	N = 0 - 1i
	E = 1 + 0i
	S = 0 + 1i
	W = -1 + 0i
)

type Grid struct {
	vals map[complex128]int
	w    int
	h    int
}

func parseInput(input string) Grid {
	grid := Grid{
		vals: map[complex128]int{},
		w:    0,
		h:    0,
	}

	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	y, w := 0.0, 0

	for scanner.Scan() {
		line := scanner.Text()
		w = len(line)

		for x, c := range line {
			v := int(c - '0')

			grid.vals[complex(float64(x), y)] = v
		}

		y++
	}

	grid.w = w
	grid.h = int(y)

	return grid
}

func (g Grid) Out(p complex128) bool {
	_, ok := g.vals[p]

	return !ok
}

func (g Grid) Visible(p complex128) bool {
	v := g.vals[p]
	comp := [4]complex128{N, E, S, W}

dirsLoop:
	for _, c := range comp {
		k := p + c
		for {
			n, ok := g.vals[k]
			if !ok {
				return true
			} else if n >= v {
				continue dirsLoop
			}
			k += c
		}
	}

	return false
}

func (g Grid) Count(p complex128) int {
	v := g.vals[p]
	comp := [4]complex128{N, E, S, W}
	scenic := 1

dirsLoop:
	for _, c := range comp {
		k := p + c
		tot := 0

		for {
			n, ok := g.vals[k]

			if !ok {
				scenic *= tot
				continue dirsLoop
			}
			tot++

			if n >= v {
				scenic *= tot
				continue dirsLoop
			}

			k += c
		}
	}

	return scenic
}

func Part1(input string) int {
	grid := parseInput(input)

	tot := 0

	for key := range grid.vals {
		if grid.Visible(key) {
			tot++
		}
	}

	return tot
}

func Part2(input string) int {
	grid := parseInput(input)

	highest := 0

	for key := range grid.vals {
		tot := grid.Count(key)

		if tot > highest {
			highest = tot
		}
	}

	return highest
}

func Run(input string) {
	fmt.Println("Day 8 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
