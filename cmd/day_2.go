package cmd

import (
	"fmt"
	"github.com/pinoniq/aoc-2024/utils"
)

func Day2Bis() {
	input := utils.ReadFile("day_2.txt")
	reportLines := utils.SplitLines(input)

	safeReports := 0

	for _, line := range reportLines {
		report := utils.SplitOnMultipleSpacesAsInt(line)

		if isSafeReport(report) {
			safeReports++
		} else {
			// try brute forcing
			for i := 0; i < len(report); i++ {
				dampenedReport := removeAtIndex(report, i)
				if isSafeReport(dampenedReport) {
					safeReports++
					break
				}
			}
		}

	}

	fmt.Println("reports: ", safeReports)
}

func removeAtIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice // return the original slice if index is out of bounds
	}

	newSlice := make([]int, 0)

	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}

func Day2() {
	input := utils.ReadFile("day_2.txt")
	reportLines := utils.SplitLines(input)

	safeReports := 0

	for _, line := range reportLines {
		report := utils.SplitOnMultipleSpacesAsInt(line)

		if isSafeReport(report) {
			safeReports++
		}

	}

	fmt.Println("reports: ", safeReports)
}

func isSafeReport(report []int) bool {
	// check if in increasing or decreasing mode
	isIncreasing := report[0] < report[1]

	for i, l := 1, len(report); i < l; i++ {
		diff := absDiffInt(report[i], report[i-1])
		if diff < 1 || diff > 3 {
			return false
		}

		if isIncreasing {
			if report[i] < report[i-1] {
				return false
			}
		} else {
			if report[i] > report[i-1] {
				return false
			}
		}
	}

	return true
}
