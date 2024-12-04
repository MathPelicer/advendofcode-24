package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	reader := bufio.NewReader(file)

	input := ""
	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			input += string(line)
		}
		if err != nil {
			break
		}
	}

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
