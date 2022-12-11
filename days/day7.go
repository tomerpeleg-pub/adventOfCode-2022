package days

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type fileStructure map[string]big.Int

func parseCommands(input string) fileStructure {
	curDir := ""
	files := fileStructure{
		"/": *big.NewInt(0),
	}

	for _, line := range strings.Split(input, "\n") {
		vals := strings.Fields(line)

		if len(vals) < 2 {
			continue
		}

		if vals[0] == "$" {
			switch vals[1] {
			case "cd":
				switch vals[2] {
				case "/":
					curDir = "/"
				case "..":
					curDir = curDir[:strings.LastIndex(curDir, "/")]
				default:
					if curDir == "/" {
						curDir += vals[2]
					} else {
						curDir += "/" + vals[2]
					}
				}
			case "ls":
			}
		} else {
			switch vals[0] {
			case "dir":
				newDir := curDir + "/" + vals[1]
				if curDir == "/" {
					newDir = "/" + vals[1]
				}
				_, ok := files[newDir]

				if !ok {
					files[newDir] = *big.NewInt(0)
				}
			default:
				size, _ := strconv.ParseInt(vals[0], 10, 64)
				bigSize := *big.NewInt(size)
				bigFile := files[curDir]
				files[curDir] = *bigSize.Add(&bigSize, &bigFile)
			}
		}
	}

	newFiles := fileStructure{}

	for dir, dirSize := range files {
		subTotal := big.NewInt(0)

		for sub, subSize := range files {
			if sub != dir && strings.Contains(sub, dir) {
				subTotal.Add(subTotal, &subSize)
			}
		}

		newFiles[dir] = *subTotal.Add(subTotal, &dirSize)
	}

	return newFiles
}

func Day7Part1(input string) int64 {
	files := parseCommands(input)

	var total big.Int = *big.NewInt(0)

	min := big.NewInt(100000)
	for _, size := range files {
		if size.Cmp(min) <= 0 {
			total = *total.Add(&size, &total)
		}
	}

	return total.Int64()
}

func Day7Part2(input string) int64 {
	totalSpace := big.NewInt(70000000)
	requiredSpace := big.NewInt(30000000)

	files := parseCommands(input)

	rootSize := files["/"]

	freeSpace := totalSpace.Sub(totalSpace, &rootSize)
	minSpace := requiredSpace.Sub(requiredSpace, freeSpace)

	smallest := rootSize

	for _, size := range files {
		if size.Cmp(minSpace) >= 0 && size.Cmp(&smallest) < 0 {
			smallest = size
		}
	}

	return smallest.Int64()
}

func Day7(input string) {
	fmt.Println("Day 7 -----")
	fmt.Println("Part 1:", Day7Part1(input))
	fmt.Println("Part 2:", Day7Part2(input))
}
