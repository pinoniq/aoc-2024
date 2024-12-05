package cmd

import (
	"fmt"
	"github.com/pinoniq/aoc-2024/utils"
	"strings"
)

func Day4() {
	input := utils.ReadFile("day_4.txt")
	inputLines := strings.Split(input, "\r\n")

	numOfLines := len(inputLines)
	numOfCols := len(inputLines[0])

	charMap := make(map[int]map[int]string, numOfLines)

	for i, line := range inputLines {
		charMap[i] = make(map[int]string, numOfCols)
		for j, char := range line {
			charMap[i][j] = string(char)
		}
	}

	// check for XMAS, always starting at the X
	numXmas := 0
	for i, line := range inputLines {
		for j, char := range line {
			if char != 'X' {
				continue
			}

			// check horizontal right
			if j+3 < numOfCols {
				if line[j:j+4] == "XMAS" {
					numXmas++
				}
			}

			// diagonal top right
			if i-3 >= 0 && j+3 < numOfCols {
				if charMap[i-1][j+1] == "M" && charMap[i-2][j+2] == "A" && charMap[i-3][j+3] == "S" {
					numXmas++
				}
			}

			// top
			if i-3 >= 0 {
				if charMap[i-1][j] == "M" && charMap[i-2][j] == "A" && charMap[i-3][j] == "S" {
					numXmas++
				}
			}

			// diagonal top left
			if j-3 >= 0 && i-3 >= 0 {
				if charMap[i-1][j-1] == "M" && charMap[i-2][j-2] == "A" && charMap[i-3][j-3] == "S" {
					numXmas++
				}
			}

			// horizontal left
			if j-3 >= 0 {
				if line[j-3:j+1] == "SAMX" {
					numXmas++
				}
			}

			// diagonal bottom left
			if i+3 < numOfLines && j-3 >= 0 {
				if charMap[i+1][j-1] == "M" && charMap[i+2][j-2] == "A" && charMap[i+3][j-3] == "S" {
					numXmas++
				}
			}

			// bottom
			if i+3 < numOfLines {
				if charMap[i+1][j] == "M" && charMap[i+2][j] == "A" && charMap[i+3][j] == "S" {
					numXmas++
				}
			}

			// diagonal bottom right
			if i+3 < numOfLines && j+3 < numOfCols {
				if charMap[i+1][j+1] == "M" && charMap[i+2][j+2] == "A" && charMap[i+3][j+3] == "S" {
					numXmas++
				}
			}
		}
	}

	// clean up invalid X
	fmt.Println("Xmas found: ", numXmas)
}

func printMap(charMap map[int]map[int]string) {
	numOfLines := len(charMap)
	numOfCols := len(charMap[0])
	for i := 0; i < numOfLines; i++ {
		for j := 0; j < numOfCols; j++ {
			fmt.Print(charMap[i][j], " ")
		}
		fmt.Println() // Newline after each row
	}
}

func Day4Bis() {
	input := utils.ReadFile("day_4.txt")
	inputLines := strings.Split(input, "\r\n")

	numOfLines := len(inputLines)
	numOfCols := len(inputLines[0])

	charMap := make(map[int]map[int]string, numOfLines)

	for i, line := range inputLines {
		charMap[i] = make(map[int]string, numOfCols)
		for j, char := range line {
			charMap[i][j] = string(char)
		}
	}

	// check for XMAS, always starting at the X
	numXmas := 0
	for i, line := range inputLines {
		// make sure there is room
		if i < 1 {
			continue
		}

		if i+1 >= numOfLines {
			continue
		}

		for j, char := range line {
			if char != 'A' {
				continue
			}

			if j < 1 {
				continue
			}

			if j+1 >= numOfCols {
				continue
			}

			// check horizontal right
			topLeft := charMap[i-1][j-1]
			topRight := charMap[i-1][j+1]
			bottomRight := charMap[i+1][j+1]
			bottomLeft := charMap[i+1][j-1]

			bottomLeftToTopRight := topLeft == "M" && bottomRight == "S" || topLeft == "S" && bottomRight == "M"
			bottomRightToTopLeft := topRight == "M" && bottomLeft == "S" || topRight == "S" && bottomLeft == "M"

			if bottomLeftToTopRight && bottomRightToTopLeft {
				numXmas++
			}
		}
	}

	// clean up invalid X
	fmt.Println("Xmas found: ", numXmas)
}
