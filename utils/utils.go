package utils

import (
	"bufio"
	"os"
)

func ReadFileIntoString(filePath string) (string, error) {
	file, _ := os.Open(filePath)

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

	return input, nil
}

func ReadFileInto2dArray(filePath string) [][]string {
	file, _ := os.Open(filePath)

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

	return inputMatrix
}
