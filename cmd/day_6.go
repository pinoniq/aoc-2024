package cmd

import (
	"fmt"
	"github.com/pinoniq/aoc-2024/utils"
	"slices"
	"strings"
)

var (
	obstacle = '#'
	guard    = '^'
)

type guardDirection int

const (
	up guardDirection = iota
	down
	left
	right
)

type guardPosition struct {
	x, y int
}

func Day6() {
	input := utils.ReadFile("day_6.txt")
	inputLines := strings.Split(input, "\r\n")

	obstacles := make(map[int]map[int]bool, len(inputLines))
	path := make(map[int]map[int][]guardDirection)

	maxX := len(inputLines)
	maxY := len(inputLines[0])

	guardCurrentPosition := guardPosition{}
	guardCurrentDirection := up

	for i, line := range inputLines {
		for j, char := range line {
			if char == obstacle {
				if _, ok := obstacles[i]; !ok {
					obstacles[i] = make(map[int]bool)
				}
				obstacles[i][j] = true
			}

			if char == guard {
				guardCurrentPosition.x = i
				guardCurrentPosition.y = j
			}
		}
	}

	guardCanContinue := true

	nextGuardPosition := guardPosition{
		x: guardCurrentPosition.x,
		y: guardCurrentPosition.y,
	}

	if _, ok := path[nextGuardPosition.x]; !ok {
		path[nextGuardPosition.x] = make(map[int][]guardDirection)
	}
	if _, ok := path[nextGuardPosition.x][nextGuardPosition.y]; !ok {
		path[nextGuardPosition.x][nextGuardPosition.y] = []guardDirection{guardCurrentDirection}
	} else {
		if !slices.Contains(path[nextGuardPosition.x][nextGuardPosition.y], guardCurrentDirection) {
			path[nextGuardPosition.x][nextGuardPosition.y] = append(path[nextGuardPosition.x][nextGuardPosition.y], guardCurrentDirection)
		}
	}

	for guardCanContinue {
		calculateNextGuardPosition(guardCurrentPosition, guardCurrentDirection, &nextGuardPosition)

		// if next position is outside of map, end
		if nextGuardPosition.x < 0 || nextGuardPosition.x >= maxX || nextGuardPosition.y < 0 || nextGuardPosition.y >= maxY {
			guardCanContinue = false
			break
		}

		// reset next guard position and rotate 90 degrees if at an obstacle
		if _, ok := obstacles[nextGuardPosition.x]; ok {
			if _, ok := obstacles[nextGuardPosition.x][nextGuardPosition.y]; ok {
				nextGuardPosition.y = guardCurrentPosition.y
				nextGuardPosition.x = guardCurrentPosition.x
				guardCurrentDirection = calculateNextGuardDirection(guardCurrentDirection)
				continue
			}
		}

		// walk to next position
		if _, ok := path[nextGuardPosition.x]; !ok {
			path[nextGuardPosition.x] = make(map[int][]guardDirection)
		}
		if _, ok := path[nextGuardPosition.x][nextGuardPosition.y]; !ok {
			path[nextGuardPosition.x][nextGuardPosition.y] = []guardDirection{guardCurrentDirection}
		} else {
			if !slices.Contains(path[nextGuardPosition.x][nextGuardPosition.y], guardCurrentDirection) {
				path[nextGuardPosition.x][nextGuardPosition.y] = append(path[nextGuardPosition.x][nextGuardPosition.y], guardCurrentDirection)
			}
		}

		guardCurrentPosition.y = nextGuardPosition.y
		guardCurrentPosition.x = nextGuardPosition.x
	}

	// clean up invalid X
	fmt.Println("path: ", path)
	count := 0
	for _, l := range path {
		for _, _ = range l {
			count++
		}
	}
	fmt.Println("count: ", count)
}

func doesLoop(guardCurrentPosition guardPosition, guardCurrentDirection guardDirection, maxX int, maxY int, obstacles map[int]map[int]bool) bool {
	path := make(map[int]map[int][]guardDirection)
	guardCanContinue := true

	nextGuardPosition := guardPosition{
		x: guardCurrentPosition.x,
		y: guardCurrentPosition.y,
	}

	if _, ok := path[nextGuardPosition.x]; !ok {
		path[nextGuardPosition.x] = make(map[int][]guardDirection)
	}
	if _, ok := path[nextGuardPosition.x][nextGuardPosition.y]; !ok {
		path[nextGuardPosition.x][nextGuardPosition.y] = []guardDirection{guardCurrentDirection}
	} else {
		if !slices.Contains(path[nextGuardPosition.x][nextGuardPosition.y], guardCurrentDirection) {
			path[nextGuardPosition.x][nextGuardPosition.y] = append(path[nextGuardPosition.x][nextGuardPosition.y], guardCurrentDirection)
		}
	}

	for guardCanContinue {
		calculateNextGuardPosition(guardCurrentPosition, guardCurrentDirection, &nextGuardPosition)

		// if next position is outside of map, end
		if nextGuardPosition.x < 0 || nextGuardPosition.x >= maxX || nextGuardPosition.y < 0 || nextGuardPosition.y >= maxY {
			guardCanContinue = false
			break
		}

		// reset next guard position and rotate 90 degrees if at an obstacle
		if _, ok := obstacles[nextGuardPosition.x]; ok {
			if _, ok := obstacles[nextGuardPosition.x][nextGuardPosition.y]; ok {
				if obstacles[nextGuardPosition.x][nextGuardPosition.y] {
					nextGuardPosition.y = guardCurrentPosition.y
					nextGuardPosition.x = guardCurrentPosition.x
					guardCurrentDirection = calculateNextGuardDirection(guardCurrentDirection)
					continue
				}
			}
		}

		// walk to next position
		if _, ok := path[nextGuardPosition.x]; !ok {
			path[nextGuardPosition.x] = make(map[int][]guardDirection)
		}
		if _, ok := path[nextGuardPosition.x][nextGuardPosition.y]; !ok {
			path[nextGuardPosition.x][nextGuardPosition.y] = []guardDirection{guardCurrentDirection}
		} else {
			if !slices.Contains(path[nextGuardPosition.x][nextGuardPosition.y], guardCurrentDirection) {
				path[nextGuardPosition.x][nextGuardPosition.y] = append(path[nextGuardPosition.x][nextGuardPosition.y], guardCurrentDirection)
			} else {
				return true
			}
		}

		guardCurrentPosition.y = nextGuardPosition.y
		guardCurrentPosition.x = nextGuardPosition.x
	}

	return false
}

func calculateNextGuardPosition(currentGuardPosition guardPosition, currentGuardDirection guardDirection, nextGuardPosition *guardPosition) {
	switch currentGuardDirection {
	case up:
		nextGuardPosition.x = currentGuardPosition.x - 1
		nextGuardPosition.y = currentGuardPosition.y
	case down:
		nextGuardPosition.x = currentGuardPosition.x + 1
		nextGuardPosition.y = currentGuardPosition.y
	case left:
		nextGuardPosition.x = currentGuardPosition.x
		nextGuardPosition.y = currentGuardPosition.y - 1
	case right:
		nextGuardPosition.x = currentGuardPosition.x
		nextGuardPosition.y = currentGuardPosition.y + 1
	}
}

func calculateNextGuardDirection(currentGuardDirection guardDirection) guardDirection {
	switch currentGuardDirection {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		return up
	}
}

func Day6Bis() {
	input := utils.ReadFile("day_6.txt")
	inputLines := strings.Split(input, "\r\n")

	obstacles := make(map[int]map[int]bool, len(inputLines))

	maxX := len(inputLines)
	maxY := len(inputLines[0])

	guardCurrentPosition := guardPosition{}
	guardCurrentDirection := up

	for i, line := range inputLines {
		for j, char := range line {
			if char == obstacle {
				if _, ok := obstacles[i]; !ok {
					obstacles[i] = make(map[int]bool)
				}
				obstacles[i][j] = true
			}

			if char == guard {
				guardCurrentPosition.x = i
				guardCurrentPosition.y = j
			}
		}
	}

	foundLoops := 0
	for i, line := range inputLines {
		for j, char := range line {
			if char != obstacle && char != guard {
				if _, ok := obstacles[i]; !ok {
					obstacles[i] = make(map[int]bool)
				}
				obstacles[i][j] = true
				if doesLoop(guardCurrentPosition, guardCurrentDirection, maxX, maxY, obstacles) {
					foundLoops++
				}
				obstacles[i][j] = false
			}
		}
	}

	fmt.Println("total loops: ", foundLoops)
}
