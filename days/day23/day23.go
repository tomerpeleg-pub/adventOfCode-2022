package day23

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	NO = -10000
)

const (
	N  = 0
	NE = 1
	E  = 2
	SE = 3
	S  = 4
	SW = 5
	W  = 6
	NW = 7
)

type Grid struct {
	rect  Rect
	elves map[complex128]bool
}

type Rect struct {
	x, y, w, h int
}

func (g Grid) Rect() Rect {
	s := false
	rect := Rect{}

	for p, v := range g.elves {
		if v {
			x := int(real(p))
			y := int(imag(p))

			if x < rect.x || !s {
				rect.x = x
				s = true
			} else if x > rect.w {
				rect.w = x
			}

			if y < rect.y || !s {
				rect.y = y
				s = true
			} else if y > rect.h {
				rect.h = y
			}
		}
	}
	rect.w++
	rect.h++
	return rect
}

func (g Grid) GetC(p complex128) bool {
	return g.elves[p]
}

func (g Grid) Out(p complex128) bool {
	x := int(real(p))
	y := int(imag(p))

	if x < g.rect.x || x > g.rect.w || y < g.rect.y || y > g.rect.h {
		return true
	}
	return false
}

func (g Grid) Get(x, y int) bool {
	return g.elves[complex(float64(x), float64(y))]
}

func (g *Grid) Set(x, y int, v bool) {
	if g.elves == nil {
		g.elves = map[complex128]bool{}
	}
	g.elves[complex(float64(x), float64(y))] = v
}

func (g Grid) String() string {
	rect := g.rect

	str := fmt.Sprintf("{x:%v,y:%v,w:%v,h:%v}\n", g.rect.x, g.rect.y, g.rect.w, g.rect.h)
	for y := rect.y; y < rect.h; y++ {
		for x := rect.x; x < rect.w; x++ {
			if g.Get(x, y) {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func parseInput(input string) Grid {
	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	grid := Grid{}

	y, w := 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		for x, char := range line {
			if char == '#' {
				grid.Set(x, y, true)
			}

			if x > w {
				w = x
			}
		}
		y++
	}
	grid.rect = Rect{
		0, 0, w + 1, y,
	}

	return grid
}

type Proposal struct {
	r     bool
	from  complex128
	to    complex128
	clash bool
}

func neighbours(g Grid, p complex128) [4]complex128 {
	nw := p - 1 - 1i
	n := p - 1i
	ne := p + 1 - 1i
	e := p + 1
	se := p + 1 + 1i
	s := p + 1i
	sw := p - 1 + 1i
	w := p - 1

	dirs := [4]complex128{NO, NO, NO, NO}
	f := false

	// if !g.Out(n) {
	if !g.GetC(nw) && !g.GetC(n) && !g.GetC(ne) {
		dirs[0] = n
	} else {
		f = true
	}
	// }

	// if !g.Out(s) {
	if !g.GetC(se) && !g.GetC(s) && !g.GetC(sw) {
		dirs[1] = s
	} else {
		f = true
	}
	// }

	// if !g.Out(w) {
	if !g.GetC(sw) && !g.GetC(w) && !g.GetC(nw) {
		dirs[2] = w
	} else {
		f = true
	}
	// }

	// if !g.Out(e) {
	if !g.GetC(ne) && !g.GetC(e) && !g.GetC(se) {
		dirs[3] = e
	} else {
		f = true
	}
	// }

	if !f {
		return [4]complex128{NO, NO, NO, NO}
	}
	return dirs
}

func Round1(g Grid, startDir int) map[complex128]Proposal {
	proposals := map[complex128]Proposal{}

elfLoop:
	for p, v := range g.elves {
		if v {
			dirs := neighbours(g, p)

			for i := 0; i < 4; i++ {
				v := (i + startDir) % 4
				dir := dirs[v]

				if dir != NO {
					prop := proposals[dir]

					if prop.r {
						prop.clash = true
					} else {
						prop.from = p
						prop.to = dir
						prop.r = true
					}

					proposals[dir] = prop
					continue elfLoop
				}
			}
		}
	}

	return proposals
}

func Round2(g *Grid, proposals map[complex128]Proposal) bool {
	m := false
	for _, prop := range proposals {
		if !prop.clash {
			m = true
			g.elves[prop.from] = false
			g.elves[prop.to] = true
		}
	}
	return m
}

func CountEmpty(g Grid) int {
	r := g.Rect()

	g.rect = r
	fmt.Println("Rect:", g)
	tot := (r.w - r.x) * (r.h - r.y)

	elves := 0

	for _, v := range g.elves {
		if v {
			elves++
		}
	}

	return tot - elves
}

func Part1(input string) int {
	start := parseInput(input)

	fmt.Println("=== START ===")
	fmt.Println(start)
	for i := 1; i <= 10; i++ {
		proposals := Round1(start, i-1)
		Round2(&start, proposals)

	}
	fmt.Printf("=== After round %v ===\n", 10)
	fmt.Println(start)

	return CountEmpty(start)
}

func Part2(input string) int {
	start := parseInput(input)

	fmt.Println("=== START ===")
	fmt.Println(start)
	for i := 1; true; i++ {
		proposals := Round1(start, i-1)
		didMove := Round2(&start, proposals)

		if !didMove {
			fmt.Printf("=== After round %v ===\n", i)
			fmt.Println(start)
			return i
		}
	}

	return CountEmpty(start)
}

func Run(input string) {
	fmt.Println("Day 23 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
