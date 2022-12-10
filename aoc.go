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
	}
}
