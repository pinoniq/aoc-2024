package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(fileName string) string {
	pwd, _ := os.Getwd()

	fileContent, err := os.ReadFile(fmt.Sprintf("%s/data/%s", pwd, fileName))
	if err != nil {
		panic(err)
	}
	return string(fileContent)
}

func SplitLines(input string) []string {
	return strings.Split(input, "\n")
}

func SplitOnMultipleSpaces(input string) []string {
	return strings.Fields(input)
}

func SplitOnMultipleSpacesAsInt(input string) []int {
	values := strings.Fields(input)

	intVals := make([]int, len(values))
	for i, v := range values {
		intVals[i], _ = strconv.Atoi(v)
	}

	return intVals
}
