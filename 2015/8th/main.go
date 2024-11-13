package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func CalculateDifference(line string, re *regexp.Regexp) (int, int) {
	accumulator := 0
	counter := 0

	indices := re.FindStringIndex(line)
	for indices != nil {
		accumulator += indices[1] - indices[0]
		counter++
		line = line[indices[1]:]
		indices = re.FindStringIndex(line)
	}
	difference := accumulator + 2

	return difference, counter
}

func main() {
	file, err := os.Open("list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	differenceBetweenLiteralsAndBytes, numberOfSpecialCharacters := 0, 0
	differenceBetweenEncodedAndRaw := 0
	scanner := bufio.NewScanner(file)

	var (
		escapeSequences  = regexp.MustCompile(`(\\\\|\\"|\\x[0-9a-fA-F]{2})`)
		escapeCharacters = regexp.MustCompile(`(\\|")`)
	)

	for scanner.Scan() {
		line := scanner.Text()
		difference, count := CalculateDifference(line, escapeSequences)
		differenceBetweenLiteralsAndBytes += difference
		numberOfSpecialCharacters += count

		difference, _ = CalculateDifference(line, escapeCharacters)
		differenceBetweenEncodedAndRaw += difference
	}

	fmt.Println(differenceBetweenLiteralsAndBytes - numberOfSpecialCharacters)
	fmt.Println(differenceBetweenEncodedAndRaw)
}
