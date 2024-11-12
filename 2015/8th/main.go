package main

import (
	"fmt"
	"regexp"
)

var (
	escapeSequences  = regexp.MustCompile(`(\\\\|\\"|\\x\w\w)`)
	escapeCharacters = regexp.MustCompile(`(\\|"|x\w\w)`)
)

func CalculateDifference(line string) int {
	accumulator := 0
	indices := escapeSequences.FindStringIndex(line)
	for indices != nil {
		// Size of escape sequence - size of ASCII character (bytes).
		accumulator += (indices[1] - indices[0]) - 1

		line = line[indices[1]:]
		indices = escapeSequences.FindStringIndex(line)
	}
	// Always starts and ands with double quotes (+ two ASCII letters).
	difference := accumulator + 2

	return difference
}

func CalculateEncodingLength(line string) int {
	accumulator := 0
	indices := escapeCharacters.FindStringIndex(line)
	for indices != nil {
		// Size of escape sequence - size of ASCII character (bytes).
		accumulator += indices[1] - indices[0] - 1

		line = line[indices[1]:]
		indices = escapeSequences.FindStringIndex(line)
	}
	// Always starts and ands with double quotes (+ two ASCII letters).
	difference := accumulator + 4

	return difference
}

func main() {
	// file, err := os.Open("list.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// differenceBetweenLiteralsAndBytes := 0
	// encodingCharactersCount := 0
	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	differenceBetweenLiteralsAndBytes += CalculateDifference(line)
	// 	encodingCharactersCount += CalculateEncodingLength(line)
	// }

	// fmt.Println(differenceBetweenLiteralsAndBytes)
	// fmt.Println(encodingCharactersCount)

	testCases := []string{
		`""`,
		`"abc"`,
		`"aaa\"aaa"`,
		`"\x27"`,
	}

	encodingCharactersDifference := 0

	for _, line := range testCases {
		encodingCharactersDifference += CalculateEncodingLength(line)
	}

	fmt.Println(encodingCharactersDifference)
}
