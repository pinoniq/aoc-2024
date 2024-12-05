package cmd

import (
	"fmt"
	"github.com/pinoniq/aoc-2024/utils"
	"sort"
	"strconv"
	"strings"
)

type pageOrderingRule struct {
	page1, page2 int
}

func (p *pageOrderingRule) valid(pages map[int]int) bool {
	page1Index, page1Exists := pages[p.page1]
	page2Index, page2Exists := pages[p.page2]

	if !page1Exists || !page2Exists {
		return true
	}

	return page1Index < page2Index
}

func (p *pageOrderingRule) applies(p1, p2 int) bool {
	return p1 == p.page1 && p2 == p.page2 || p1 == p.page2 && p2 == p.page1
}

func (p *pageOrderingRule) validate(p1, p2 int) bool {
	return p1 == p.page1 && p2 == p.page2
}

func Day5() {
	input := utils.ReadFile("day_5.txt")
	inputLines := strings.Split(input, "\r\n\r\n")

	pageOrderingRules := make([]pageOrderingRule, len(inputLines[0]))

	for i, pageOrderingRuleLine := range strings.Split(inputLines[0], "\r\n") {
		pagesInRule := strings.Split(pageOrderingRuleLine, "|")

		page1, _ := strconv.Atoi(pagesInRule[0])
		page2, _ := strconv.Atoi(pagesInRule[1])

		pageOrderingRules[i] = pageOrderingRule{
			page1: page1,
			page2: page2,
		}
	}

	validPagesMiddleSum := 0
	for _, pagesLine := range strings.Split(inputLines[1], "\r\n") {
		pagesString := strings.Split(pagesLine, ",")
		pageNumbersMap := make(map[int]int)
		var middleNumber int
		for i, pageString := range pagesString {
			page, _ := strconv.Atoi(pageString)
			pageNumbersMap[page] = i
			if i == len(pagesString)/2 {
				middleNumber = page
			}
		}

		if validatePages(pageNumbersMap, pageOrderingRules) {
			validPagesMiddleSum += middleNumber
		}
	}

	// clean up invalid X
	fmt.Println("validPagesMiddleSum: ", validPagesMiddleSum)
}

func validatePages(pageNumbers map[int]int, pageOrderingRules []pageOrderingRule) bool {
	for _, pageOrderingRule := range pageOrderingRules {
		if !pageOrderingRule.valid(pageNumbers) {
			return false
		}
	}

	return true
}

func Day5Bis() {
	input := utils.ReadFile("day_5.txt")
	inputLines := strings.Split(input, "\r\n\r\n")

	pageOrderingRules := make([]pageOrderingRule, len(inputLines[0]))

	for i, pageOrderingRuleLine := range strings.Split(inputLines[0], "\r\n") {
		pagesInRule := strings.Split(pageOrderingRuleLine, "|")

		page1, _ := strconv.Atoi(pagesInRule[0])
		page2, _ := strconv.Atoi(pagesInRule[1])

		pageOrderingRules[i] = pageOrderingRule{
			page1: page1,
			page2: page2,
		}
	}

	validPagesMiddleSum := 0
	for _, pagesLine := range strings.Split(inputLines[1], "\r\n") {
		pagesString := strings.Split(pagesLine, ",")
		pageNumbersMap := make(map[int]int)
		pageNumbers := make([]int, len(pagesString))
		for i, pageString := range pagesString {
			page, _ := strconv.Atoi(pageString)
			pageNumbers[i] = page
			pageNumbersMap[page] = i
		}

		// sort them according to our own ordering rules
		// we could use a map of slices, but let's just brute force it. go should be fast enough
		if !validatePages(pageNumbersMap, pageOrderingRules) {
			sort.Slice(pageNumbers, func(i, j int) bool {
				for _, pageOrderingRule := range pageOrderingRules {
					if pageOrderingRule.applies(pageNumbers[i], pageNumbers[j]) {
						return pageOrderingRule.validate(pageNumbers[i], pageNumbers[j])
					}
				}
				fmt.Println("unfound rule", pageNumbers[i], pageNumbers[j])
				return false
			})
			validPagesMiddleSum += pageNumbers[len(pageNumbers)/2]
		}
	}

	// clean up invalid X
	fmt.Println("validPagesMiddleSum: ", validPagesMiddleSum)
}
