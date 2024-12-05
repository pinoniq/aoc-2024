package cmd

import (
	"fmt"
	"github.com/pinoniq/aoc-2024/utils"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	input := utils.ReadFile("day_3.txt")

	total := 0

	for _, mul := range getMuls(input) {
		mulComponents := extractDigits(mul)
		total += mulComponents[0] * mulComponents[1]
	}

	fmt.Println("Sum:", total)
}

func extractDigits(line string) []int {
	// remove mul( and )
	digitString := line[4 : len(line)-1]

	digits := strings.Split(digitString, ",")

	intDigits := make([]int, len(digits))
	for i, d := range digits {
		intDigits[i], _ = strconv.Atoi(d)
	}

	return intDigits
}

func getMuls(line string) []string {
	pattern := `mul\(\d+,\d+\)` // Replace with your actual regex pattern
	re := regexp.MustCompile(pattern)

	return re.FindAllString(line, -1)
}

func Day3Bis() {
	input := utils.ReadFile("day_3.txt")

	total := 0

	splitOnDos := strings.Split(input, "do()")

	for _, line := range splitOnDos {
		splitOnDonts := strings.Split(line, "don't()")

		for _, mul := range getMuls(splitOnDonts[0]) {
			mulComponents := extractDigits(mul)
			total += mulComponents[0] * mulComponents[1]
		}
	}

	fmt.Println("Sum:", total)
}
