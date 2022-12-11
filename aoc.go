package main

import (
	"log"
	"os"

	"github.com/tomerpeleg-pub/aoc2022/days"
)

func GetDayInput(day string) string {

	content, err := os.ReadFile("inputs/day" + day)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func RunDay(day string) {
	input := GetDayInput(day)

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
	case "7":
		days.Day7(input)
	}
}

func main() {
	day := os.Args[1]

	RunDay(day)
}
