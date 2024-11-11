package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IsNice(input string) bool {
	if len(input) < 3 {
		return false
	}

	vowelsCount := 0
	twoEqualConsecutive := false

	switch input[0] {
	case 'a', 'e', 'i', 'o', 'u':
		vowelsCount++
	}

	for i := 1; i < len(input); i++ {
		// First condition - at least 3 vowels.
		letter := input[i]
		switch letter {
		case 'a', 'e', 'i', 'o', 'u':
			vowelsCount++
		}

		// Second condition - two same consecutive letters.
		if input[i-1] == letter {
			twoEqualConsecutive = true
		}

		// Third condition - banned combinations.
		switch input[i-1 : i+1] {
		case "ab", "cd", "pq", "xy":
			return false
		}
	}

	return vowelsCount >= 3 && twoEqualConsecutive
}

func IsNiceNewRules(input string) bool {
	if len(input) < 3 {
		return false
	}

	twoLettersPattern := false
	letterInMiddle := false

	for i := 0; i < len(input)-2; i++ {
		// First condition.
		if !twoLettersPattern && strings.Contains(input[i+2:], input[i:i+2]) {
			twoLettersPattern = true
		}

		// Second condition.
		if !letterInMiddle && input[i] == input[i+2] {
			letterInMiddle = true
		}

		if twoLettersPattern && letterInMiddle {
			return true
		}
	}

	return false
}

func main() {
	now := time.Now()

	defer func(startTime time.Time) {
		fmt.Printf("Time elapsed: %v\n", time.Since(startTime))
	}(now)

	countNiceStrings := func(inputFile string, isNice func(string) bool) int {
		file, err := os.Open(inputFile)
		check(err)
		defer file.Close()

		niceStringCount := 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if isNice(line) {
				niceStringCount += 1
			}
		}

		check(scanner.Err())

		return niceStringCount
	}

	fmt.Println("Number of nice strings: ", countNiceStrings("input.txt", IsNice))
	fmt.Println("Number of nice strings: ", countNiceStrings("input.txt", IsNiceNewRules))
}
