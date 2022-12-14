package main

import (
	"testing"

	"github.com/tomerpeleg-pub/aoc2022/days/day6"
	"github.com/tomerpeleg-pub/aoc2022/util"
)

func BenchmarkAOC2(b *testing.B) {
	input := util.GetDayInput("6")

	for i := 0; i < b.N; i++ {
		day6.Day6Part1(input)
	}
}

func BenchmarkAOC3(b *testing.B) {
	input := util.GetDayInput("6")

	for i := 0; i < b.N; i++ {
		day6.Day6Part2(input)
	}
}
