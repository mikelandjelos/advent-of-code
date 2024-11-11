package main

import (
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	instructions, err := os.ReadFile("directions.txt")
	check(err)

	santaCurrentLocation := Point{0, 0}
	roboSantaCurrentLocation := Point{0, 0}
	visitedLocations := map[Point]int{
		{0, 0}: 2,
	}

	locationChangeOperator := func(instruction byte, currentPosition Point) Point {
		var newLocation Point
		switch instruction {
		case '>':
			newLocation = Point{currentPosition.X + 1, currentPosition.Y}
		case '<':
			newLocation = Point{currentPosition.X - 1, currentPosition.Y}
		case '^':
			newLocation = Point{currentPosition.X, currentPosition.Y + 1}
		case 'v':
			newLocation = Point{currentPosition.X, currentPosition.Y - 1}
		default:
			panic(fmt.Errorf("unknown instruction %v", instruction))
		}
		return newLocation
	}

	uniqueLocationsVisited := 1

	for i := 0; i < len(instructions)-1; i += 2 {
		santaCurrentLocation = locationChangeOperator(instructions[i], santaCurrentLocation)
		roboSantaCurrentLocation = locationChangeOperator(instructions[i+1], roboSantaCurrentLocation)

		if _, visited := visitedLocations[santaCurrentLocation]; !visited {
			uniqueLocationsVisited += 1
			visitedLocations[santaCurrentLocation] = 0
		}
		visitedLocations[santaCurrentLocation]++

		if _, visited := visitedLocations[roboSantaCurrentLocation]; !visited {
			uniqueLocationsVisited += 1
			visitedLocations[roboSantaCurrentLocation] = 0
		}

		visitedLocations[roboSantaCurrentLocation]++
	}

	fmt.Println("Unique locations visited: ", uniqueLocationsVisited)
}
