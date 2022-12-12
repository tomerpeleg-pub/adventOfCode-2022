package main

import (
	"log"
	"os"

	"github.com/tomerpeleg-pub/aoc2022/days/day1"
	"github.com/tomerpeleg-pub/aoc2022/days/day2"
	"github.com/tomerpeleg-pub/aoc2022/days/day3"
	"github.com/tomerpeleg-pub/aoc2022/days/day4"
	"github.com/tomerpeleg-pub/aoc2022/days/day5"
	"github.com/tomerpeleg-pub/aoc2022/days/day6"
	"github.com/tomerpeleg-pub/aoc2022/days/day7"
	"github.com/tomerpeleg-pub/aoc2022/days/day8"
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
		day1.Day1(input)
	case "2":
		day2.Day2(input)
	case "3":
		day3.Day3(input)
	case "4":
		day4.Day4(input)
	case "5":
		day5.Day5(input)
	case "6":
		day6.Day6(input)
	case "7":
		day7.Day7(input)
	case "8":
		day8.Run(input)
	}
}

func main() {
	day := os.Args[1]

	RunDay(day)
}
