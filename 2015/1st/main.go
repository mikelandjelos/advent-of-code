package main

import (
	"fmt"
	"os"
)

func FindFloor(instructions []uint8) int {
	floor := 0

	for _, instruction := range instructions {
		switch instruction {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	return floor
}

func FindBasementEnteringInstruction(instructions []uint8) uint {
	floor, position := 0, uint(1)

	for _, instruction := range instructions {
		switch instruction {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			break
		}

		position++
	}

	return position
}

func main() {
	instructions, err := os.ReadFile("instructions.txt")

	if err != nil {
		panic(err)
	}

	floorReached := FindFloor(instructions)
	basementEnteringInstruction := FindBasementEnteringInstruction(instructions)

	fmt.Println(floorReached, basementEnteringInstruction)
}
