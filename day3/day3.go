package main

import (
	"advent-of-code-24/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input, _ := utils.ReadFileIntoString("input.txt")
	mulPatterns := FindAllFromRegex("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don't\\(\\)", input)

	sum := 0
	isEnabled := true
	for index, rxString := range mulPatterns {
		if rxString == "don't()" {
			isEnabled = false
			fmt.Printf("dont't() - %d\n", index)
			continue
		}
		if rxString == "do()" {
			fmt.Printf("do() - %d\n", index)
			isEnabled = true
			continue
		}

		if isEnabled == true {
			fmt.Printf("enabled mul() - %d\n", index)
			numbers := FindAllFromRegex("\\d+", rxString)
			number1, _ := strconv.Atoi(numbers[0])
			number2, _ := strconv.Atoi(numbers[1])
			sum += number1 * number2
		} else {
			continue
		}
	}

	fmt.Print(sum)
}

func FindAllFromRegex(regexExp string, text string) []string {
	regex, _ := regexp.Compile(regexExp)
	patterns := regex.FindAllString(text, -1)

	return patterns
}
