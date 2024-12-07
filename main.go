package main

import (
	"fmt"
	"github.com/pinoniq/aoc-2024/cmd"
	"time"
)

func main() {
	start := time.Now()

	// cmd.Day1()
	// cmd.Day1Bis()
	// cmd.Day2()
	// cmd.Day2Bis()
	// cmd.Day3()
	// cmd.Day3Bis()
	// cmd.Day4()
	// cmd.Day4Bis()
	// cmd.Day5()
	// cmd.Day5Bis()
	// cmd.Day6()
	// cmd.Day6Bis()
	// cmd.Day7()
	cmd.Day7Bis()

	duration := time.Since(start)
	fmt.Println(fmt.Sprintf("Time spent: %d ms", duration.Milliseconds()))
}
