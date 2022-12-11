package main

import (
	"testing"

	"github.com/tomerpeleg-pub/aoc2022/days"
)

func BenchmarkAOC2(b *testing.B) {
	input := GetDayInput("6")

	for i := 0; i < b.N; i++ {
		days.Day6Part1(input)
	}
}

func BenchmarkAOC3(b *testing.B) {
	input := GetDayInput("6")

	for i := 0; i < b.N; i++ {
		days.Day6Part2(input)
	}
}
