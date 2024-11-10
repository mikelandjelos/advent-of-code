package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseDimensions(dimensionsLabel string) (l, w, h int) {
	strings.Split(dimensionsLabel, "x")
	panic("NotImplemented")
	return dimensions[:]
}

func calculateSingleBoxResources(l, w, h int) (surface, slack int) {
	lw, wh, hl := l*w, w*h, h*l
	surface = (lw + wh + hl) << 1
	slack = min(lw, wh, hl)
	return
}

func CalculateResourcesNeeded(inputFile string) (int, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	accumulator := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return 0, nil
}

func main() {

	resourcesNeeded := CalculateResourcesNeeded("example.txt")
}
