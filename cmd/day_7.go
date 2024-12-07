package cmd

import (
	"fmt"
	"github.com/pinoniq/aoc-2024/utils"
	"strconv"
	"strings"
)

func Day7() {
	input := utils.ReadFile("day_7.txt")
	inputLines := strings.Split(input, "\r\n")

	sum := 0

	for _, line := range inputLines {
		lineParts := strings.Split(line, ":")
		equation, _ := strconv.Atoi(lineParts[0])
		numbers := utils.SplitOnMultipleSpacesAsInt(lineParts[1])

		if doesCalculate(numbers, equation, false) {
			sum += equation
		}

	}

	fmt.Println("count: ", sum)
}

func doesCalculate(remaining []int, num int, tryConcat bool) bool {
	if len(remaining) == 1 {
		return remaining[0] == num
	}

	newRemaining := remaining[0 : len(remaining)-1]

	lastNum := remaining[len(remaining)-1]

	if doesCalculate(newRemaining, num-lastNum, tryConcat) {
		return true
	}

	if num%lastNum == 0 {
		if doesCalculate(newRemaining, num/lastNum, tryConcat) {
			return true
		}
	}

	if !tryConcat {
		return false
	}

	lastNumS := strconv.Itoa(lastNum)
	numS := strconv.Itoa(num)

	if strings.HasSuffix(numS, lastNumS) {
		newNumS := strings.TrimSuffix(numS, lastNumS)
		newNum, _ := strconv.Atoi(newNumS)
		return doesCalculate(newRemaining, newNum, tryConcat)
	}

	return false
}

func Day7Bis() {
	input := utils.ReadFile("day_7.txt")
	inputLines := strings.Split(input, "\r\n")

	sum := 0

	for _, line := range inputLines {
		lineParts := strings.Split(line, ":")
		equation, _ := strconv.Atoi(lineParts[0])
		numbers := utils.SplitOnMultipleSpacesAsInt(lineParts[1])

		if doesCalculate(numbers, equation, true) {
			sum += equation
		}

	}

	fmt.Println("count: ", sum)
}
