package day7

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type file struct {
	size int
	path []string
}

type fileStructure map[string]int

func (fs fileStructure) AddSizes(path []string, size int) fileStructure {
	r := "/"
	for _, dir := range path {
		if dir != "/" {
			r += "/" + dir
		}
		fs[r] += size
	}
	return fs
}

func parseCommands(input string) fileStructure {
	scanner := bufio.NewScanner(strings.NewReader(input))

	files := fileStructure{}
	curPath := []string{"/"}
	curSize := 0

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")

		// Ignore empty lines
		if len(args) < 2 {
			continue
		}

		switch args[1] {
		case "cd":
			// Save size of current directory to map
			files = files.AddSizes(curPath, curSize)
			curSize = 0

			// When changing directory...
			switch args[2] {
			case "/":
				// only happens once, so can ignore
				continue
			case "..":
				// When going to parent, add current directories
				// total size to all the parents'
				curPath = curPath[:len(curPath)-1]
			default:
				// When going deeper, just reset the counters
				curSize = 0
				curPath = append(curPath, args[2])
			}

		case "ls":
			// Don't care
			continue

		default:
			// Don't care about directories in ls
			if args[0] == "dir" {
				continue
			}

			// Add file size to iterator
			size, _ := strconv.Atoi(args[0])
			curSize += size
		}
	}

	// Have to repeat for the last directory
	return files.AddSizes(curPath, curSize)
}

func Day7Part1(input string) int {
	files := parseCommands(input)
	total := 0

	max := 100000
	for _, size := range files {
		if size <= max {
			total += size
		}
	}

	return total
}

func Day7Part2(input string) int {
	files := parseCommands(input)

	totalDiskSpace := 70000000
	requiredUnusedSpace := 30000000
	unusedSpace := totalDiskSpace - files["/"]
	target := requiredUnusedSpace - unusedSpace

	curMin := files["/"]

	for _, space := range files {
		if space >= target && space < curMin {
			curMin = space
		}
	}

	return curMin
}

func Day7(input string) {
	fmt.Println("Day 7 -----")
	fmt.Println("Part 1:", Day7Part1(input))
	fmt.Println("Part 2:", Day7Part2(input))
}
