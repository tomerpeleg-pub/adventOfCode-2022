package day9

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/tomerpeleg-pub/aoc2022/util"
)

var dirs map[string]complex128 = map[string]complex128{
	"L": 1,
	"R": -1,
	"U": -1i,
	"D": 1i,
}

type Move struct {
	dir complex128
	num int
}

func parseMove(input string) Move {
	vals := strings.Fields(input)
	n, _ := strconv.Atoi(vals[1])

	return Move{
		dir: dirs[vals[0]],
		num: n,
	}
}

func parseInput(input string) []Move {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	moves := []Move{}

	// Loop through each line
	for scanner.Scan() {
		moves = append(moves, parseMove(scanner.Text()))
	}

	return moves
}

func Norm(num complex128) complex128 {
	r := 0.0
	c := 0.0

	if real(num) > 0 {
		r = 1
	} else if real(num) < 0 {
		r = -1
	}
	if imag(num) > 0 {
		c = 1
	} else if imag(num) < 0 {
		c = -1
	}

	return complex(r, c)
}

func Snake(size int, moves []Move) int {

	// for each part of snake, real part is x, imag part is y
	snake := make([]complex128, size)

	// keep track of the points seen by the tail
	seen := map[complex128]bool{
		(0 + 0i): true,
	}

	for _, move := range moves {

		// for each move, perform each step invidually
		for i := 0; i < move.num; i++ {

			// move the head
			snake[0] += move.dir

			// move each of the tail pieces
			for t := range snake[1:] {

				// get the distance between tail piece and the one
				// in front of it
				dist := snake[t] - snake[t+1]
				abs := complex(math.Abs(real(dist)), math.Abs(imag(dist)))

				if real(abs) > 1 || imag(abs) > 1 {

					// move tail piece by 1 in the right direction
					snake[t+1] += Norm(dist)

					// mark as seen if the piece is the last tail
					if t == size-2 {
						seen[snake[t+1]] = true
					}
				}
			}
		}
	}

	return len(seen)
}

func Part1(input string) int {
	moves := parseInput(input)

	return Snake(2, moves)
}

func Part2(input string) int {
	moves := parseInput(input)

	return Snake(10, moves)
}

func Run(input string) {
	fmt.Println("Day 9 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}

func WithInput() {
	input := util.GetDayInput("9")
	fmt.Println("Day 9 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
