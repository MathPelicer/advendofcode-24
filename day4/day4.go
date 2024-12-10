package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	reader := bufio.NewReader(file)

	input := ""
	var inputMatrix [][]string
	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			var row []string
			for _, letter := range string(line) {
				row = append(row, string(letter))
			}
			input += string(line)
			inputMatrix = append(inputMatrix, row)
		}
		if err != nil {
			break
		}
	}

	wordToBeFound := "XMAS"

	sum := 0
	for index, line := range inputMatrix {
		for indexLetter, letter := range line {
			if letter == "X" {
				if indexLetter+3 < len(line) {
					slice := line[indexLetter+1 : indexLetter+4]
					if strings.ReplaceAll(strings.Join(slice, " "), " ", "") == "MAS" {
						sum++
					}
				}
				if indexLetter-3 >= 0 {
					slice := line[indexLetter-3 : indexLetter]
					if strings.ReplaceAll(strings.Join(slice, " "), " ", "") == "SAM" {
						sum++
					}
				}
			}

			if letter == "X" && index+3 < len(inputMatrix) {
				for wordIndex := 1; wordIndex <= len(wordToBeFound)-1; wordIndex++ {
					if string(wordToBeFound[wordIndex]) != inputMatrix[index+wordIndex][indexLetter] {
						break
					}
					if wordIndex == len(wordToBeFound)-1 {
						sum++
						fmt.Printf("FOUND VERTICAL LINE: %d - INDEX %d\n", index, indexLetter)
					}
				}
			}
			if letter == "X" && index-3 >= 0 {
				for wordIndex := 1; wordIndex <= len(wordToBeFound)-1; wordIndex++ {
					if string(wordToBeFound[wordIndex]) != inputMatrix[index-wordIndex][indexLetter] {
						break
					}
					if wordIndex == len(wordToBeFound)-1 {
						sum++
						fmt.Printf("FOUND inverse vertical LINE: %d - INDEX %d\n", index, indexLetter)
					}
				}
			}
			if letter == "X" && index-3 >= 0 && indexLetter-3 >= 0 {
				for wordIndex := 1; wordIndex <= len(wordToBeFound)-1; wordIndex++ {
					if string(wordToBeFound[wordIndex]) != inputMatrix[index-wordIndex][indexLetter-wordIndex] {
						break
					}
					if wordIndex == len(wordToBeFound)-1 {
						sum++
						fmt.Printf("FOUND DIAGONAL - - LINE: %d - INDEX %d\n", index, indexLetter)
					}
				}
			}
			if letter == "X" && index+3 < len(inputMatrix) && indexLetter+3 < len(line) {
				for wordIndex := 1; wordIndex <= len(wordToBeFound)-1; wordIndex++ {
					if string(wordToBeFound[wordIndex]) != inputMatrix[index+wordIndex][indexLetter+wordIndex] {
						break
					}
					if wordIndex == len(wordToBeFound)-1 {
						sum++
						fmt.Printf("FOUND DIAGONAL + + LINE: %d - INDEX %d\n", index, indexLetter)
					}
				}
			}
			if letter == "X" && index-3 >= 0 && indexLetter+3 < len(line) {
				for wordIndex := 1; wordIndex <= len(wordToBeFound)-1; wordIndex++ {
					if string(wordToBeFound[wordIndex]) != inputMatrix[index-wordIndex][indexLetter+wordIndex] {
						break
					}
					if wordIndex == len(wordToBeFound)-1 {
						sum++
						fmt.Printf("FOUND DIAGONAL - + LINE: %d - INDEX %d\n", index, indexLetter)
					}
				}
			}
			if letter == "X" && index+3 < len(inputMatrix) && indexLetter-3 >= 0 {
				for wordIndex := 1; wordIndex <= len(wordToBeFound)-1; wordIndex++ {
					if string(wordToBeFound[wordIndex]) != inputMatrix[index+wordIndex][indexLetter-wordIndex] {
						break
					}
					if wordIndex == len(wordToBeFound)-1 {
						sum++
						fmt.Printf("FOUND DIAGONAL + - LINE: %d - INDEX %d\n", index, indexLetter)
					}
				}
			}
		}
	}

	fmt.Printf("SUM: %d", sum)
	//sum = 2613
}
