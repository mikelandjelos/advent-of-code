package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func parseDimensions(dimensionsLabel string) (l, w, h int) {
	separateDimensions := strings.Split(dimensionsLabel, "x")

	l, err := strconv.Atoi(separateDimensions[0])
	check(err)

	w, err = strconv.Atoi(separateDimensions[1])
	check(err)

	h, err = strconv.Atoi(separateDimensions[2])
	check(err)

	return
}

func main() {

	// Opening the input file.
	file, err := os.Open("dimensions.txt")
	check(err)
	defer file.Close()

	// Main calculation.
	wrappingPaperSurface, ribbonLength := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l, w, h := parseDimensions(line)

		// Ribbon length = Smallest Side Perimeter + Volume
		ribbonLength += (min(l+w, w+h, h+l) << 1) + l*w*h

		// Wrapping paper surface = Box surface + Slack (Smallest Side Area)
		lw, wh, hl := l*w, w*h, h*l
		wrappingPaperSurface += ((lw + wh + hl) << 1) + min(lw, wh, hl)
	}

	// Checking if there was error while reading the file.
	check(scanner.Err())

	fmt.Printf("Wrapping paper length: %v [feet^2],\nRibbon length: %v [feet]", wrappingPaperSurface, ribbonLength)
}
