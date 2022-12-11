package main

import (
	"log"
	"os"

	"github.com/tomerpeleg-pub/aoc2022/days"
)

func main() {
	day := os.Args[1]

	content, err := os.ReadFile("inputs/day" + day)

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	switch day {
	case "1":
		days.Day1(input)
	case "2":
		days.Day2(input)
	case "3":
		days.Day3(input)
	case "4":
		days.Day4(input)
	case "5":
		days.Day5(input)
	case "6":
		days.Day6(input)
	}
}
