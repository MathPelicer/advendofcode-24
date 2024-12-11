package main

import (
	"advent-of-code-24/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputString, _ := utils.ReadFileIntoString("input.txt")
	inputArray := strings.Split(inputString, "\n")

	var pageOrderRules []string
	var pageUpdates []string

	for _, line := range inputArray {
		if strings.Contains(line, "|") {
			pageOrderRules = append(pageOrderRules, line)
		}

		if strings.Contains(line, ",") {
			pageUpdates = append(pageUpdates, line)
		}
	}

	orderRulesMap := make(map[string][]string)
	for _, rule := range pageOrderRules {
		ruleValues := strings.Split(rule, "|")
		orderRulesMap[ruleValues[0]] = append(orderRulesMap[ruleValues[0]], ruleValues[1])
	}

	//var isUpdateOrderValid []bool
	sumOfMiddleNumbers := 0
	for _, pageUpdate := range pageUpdates {
		update := strings.Split(pageUpdate, ",")
		var passedPages []string
		isValid := true
		for _, page := range update {
			pageRule := orderRulesMap[page]
			passedPages = append(passedPages, page)

			for _, passedPage := range passedPages {
				if slices.Contains(pageRule, passedPage) {
					isValid = false
				}
			}
		}

		if isValid == true {
			middleNumber, _ := strconv.Atoi(update[len(update)/2])
			sumOfMiddleNumbers += middleNumber
		}
	}

	fmt.Print(sumOfMiddleNumbers)
}
