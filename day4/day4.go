package main

import (
	"advent-of-code-24/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	inputMatrix := utils.ReadFileInto2dArray("input.txt")

	wordToBeFound := "XMAS"
	xMAS := []string{"MS", "SM"}

	sum := 0
	sum2 := 0
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

			if letter == "A" && index+1 < len(inputMatrix) && index-1 >= 0 && indexLetter-1 >= 0 && indexLetter+1 < len(line) {

				fmt.Printf("A at- LINE: %d - INDEX %d\n", index, indexLetter)

				firstDiagonal := inputMatrix[index-1][indexLetter-1] + inputMatrix[index+1][indexLetter+1]
				secondDiagonal := inputMatrix[index+1][indexLetter-1] + inputMatrix[index-1][indexLetter+1]

				if slices.Contains(xMAS, firstDiagonal) && slices.Contains(xMAS, secondDiagonal) {
					fmt.Printf("Found x-MAS:\n firstDiagonal %s\n secondDiagonal: %s\n", firstDiagonal, secondDiagonal)
					sum2++
				}
			}
		}
	}

	fmt.Printf("SUM: %d\n", sum)
	fmt.Printf("SUM 2: %d", sum2)
	//sum = 2613
}
