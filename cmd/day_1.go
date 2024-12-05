package cmd

import (
	"fmt"
	"github.com/pinoniq/aoc/utils"
	"sort"
)

func Day1() {
	input := utils.ReadFile("day_1.txt")
	inputLines := utils.SplitLines(input)

	leftList := make([]int, len(inputLines))
	rightList := make([]int, len(inputLines))

	for i, line := range inputLines {
		lineValues := utils.SplitOnMultipleSpacesAsInt(line)
		leftList[i] = lineValues[0]
		rightList[i] = lineValues[1]
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0

	for i := 0; i < len(leftList); i++ {
		diff := absDiffInt(leftList[i], rightList[i])
		totalDistance += diff
	}

	fmt.Printf("Day 1: %d\n", totalDistance)
}

func Day1Bis() {
	input := utils.ReadFile("day_1.txt")
	inputLines := utils.SplitLines(input)

	leftList := make([]int, len(inputLines))
	rightList := make([]int, len(inputLines))

	for i, line := range inputLines {
		lineValues := utils.SplitOnMultipleSpacesAsInt(line)
		leftList[i] = lineValues[0]
		rightList[i] = lineValues[1]
	}

	sort.Ints(rightList)

	occurrences := make(map[int]int)
	for _, value := range rightList {
		occurrences[value]++
	}

	totalSimilarity := 0

	for i := 0; i < len(leftList); i++ {
		rightListOccurence, ok := occurrences[leftList[i]]
		if ok {
			totalSimilarity += leftList[i] * rightListOccurence
		}
	}

	fmt.Printf("Day 1: %d\n", totalSimilarity)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
