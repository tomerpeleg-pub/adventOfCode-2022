package day19

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/queue"
)

type Robot struct {
	Type string
	Cost map[string]int
}

type Blueprint map[string]Robot

func parseInput(input string) []Blueprint {
	blueprintRegex := regexp.MustCompile(`Each (\w+) robot costs (\d+) (\w+)( and (\d+) (\w+))?.`)

	reader := strings.NewReader(strings.TrimSpace(input))
	scanner := bufio.NewScanner(reader)

	blueprints := []Blueprint{}

	for scanner.Scan() {
		line := scanner.Text()
		match := blueprintRegex.FindAllStringSubmatch(line, -1)

		blueprint := Blueprint{}

		for _, result := range match {
			robot := Robot{
				Type: result[1],
				Cost: map[string]int{},
			}
			primary, _ := strconv.Atoi(result[2])
			robot.Cost[result[3]] = primary

			if result[5] != "" {
				secondary, _ := strconv.Atoi(result[5])
				robot.Cost[result[6]] = secondary
			}
			blueprint[robot.Type] = robot
		}

		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}

func runBlueprint(blueprint Blueprint) {
	robots := map[string]int{"ore": 1}
	stock := map[string]int{}
	for k, v := range stock {

	}

	for min := 1; min <= 24; min++ {

		// harvest
		for material, num := range robots {
			queue.New()
		}
	}

}

func Part1(input string) int {
	blueprints := parseInput(input)

	for _, blueprint := range blueprints {
		blueprintMaximiser(blueprint)
	}

	return 12
}

func Part2(input string) int {
	return 12
}

func Run(input string) {
	fmt.Println("Day 19 -----")
	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
