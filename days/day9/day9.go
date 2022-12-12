package day9

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/tomerpeleg-pub/aoc2022/util"
)

type Point struct {
	x int
	y int
}
type Move Point

func (p Point) String() string {
	return fmt.Sprintf("%v:%v", p.x, p.y)
}

func (p Point) Move(m Move) Point {
	return Point{
		x: p.x + m.x,
		y: p.y + m.y,
	}
}

func (p Point) Catch(h Point) (Point, bool) {
	diff, abs := h.Diff(p)

	if !(abs.x > 1 || abs.y > 1) {
		return p, false
	}

	if abs.x > 0 && abs.y > 0 {
		p.x += diff.x / abs.x
		p.y += diff.y / abs.y
	} else if abs.x > 1 {
		p.x += diff.x / abs.x
	} else if abs.y > 1 {
		p.y += diff.y / abs.y
	}

	return p, abs.x > 1 || abs.y > 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p Point) Diff(p2 Point) (Point, Point) {
	return Point{
			x: p.x - p2.x,
			y: p.y - p2.y,
		}, Point{
			x: Abs(p.x - p2.x),
			y: Abs(p.y - p2.y),
		}
}

func parseMove(input string) Move {
	vals := strings.Split(input, " ")
	dir := vals[0]
	num, _ := strconv.Atoi(vals[1])

	move := Move{}

	switch dir {
	case "R":
		move.x = num
	case "U":
		move.y = -num
	case "L":
		move.x = -num
	case "D":
		move.y = num
	}

	return move
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

func Part1(input string) int {
	moves := parseInput(input)

	head := Point{x: 0, y: 0}
	tail := Point{x: 0, y: 0}

	visited := map[string]bool{"0:0": true}

	for _, move := range moves {
		head = head.Move(move)

		for t, caught := tail.Catch(head); caught; t, caught = tail.Catch(head) {
			tail = t
			visited[tail.String()] = true
		}
	}

	return len(visited)
}

func PrintSnake(snake [10]Point) {
	str := [40]string{}

	for y := 0; y < 40; y++ {

		str[y] = ""

	cell:
		for x := 0; x < 40; x++ {

			for i, t := range snake {
				if t.x == x-20 && t.y == y-20 {
					str[y] += fmt.Sprint(i)
					continue cell
				}
			}

			str[y] += "."
		}

		str[y] += "\n"
	}

	fmt.Println(str)
}

func Part2(input string) int {
	moves := parseInput(input)
	snake := [10]Point{}
	visited := map[string]bool{"0:0": true}

	for _, move := range moves {

	tails:
		for i := range snake {
			if i == 0 {
				snake[0] = snake[0].Move(move)
				continue tails
			}

			for t, more := snake[i].Catch(snake[i-1]); more; t, more = snake[i].Catch(snake[i-1]) {
				snake[i] = t

				if i == len(snake)-1 {
					visited[t.String()] = true
				}
			}
		}

		// fmt.Println("==  ", move, "  ==")
		// PrintSnake(snake)
	}

	return len(visited)
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
