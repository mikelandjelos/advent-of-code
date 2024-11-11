package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Coordinates struct {
	Row, Column int
}

type CoordinatesRange struct {
	UpperLeft  Coordinates
	LowerRight Coordinates
}

const (
	Toggle  = "toggle"
	TurnOn  = "turn on"
	TurnOff = "turn off"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var instructionCommonFormat = regexp.MustCompile(`(toggle|turn on|turn off) (\d+),(\d+) through (\d+),(\d+)`)

func parseInstruction(instruction string) (command string, through CoordinatesRange, err error) {
	matches := instructionCommonFormat.FindStringSubmatch(instruction)

	if len(matches) == 0 {
		return "", CoordinatesRange{},
			fmt.Errorf("instruction `%v` doesn't follow a common format", instruction)
	}

	// Parse the command and coordinates from the matched groups
	command = matches[1]

	upperLeftRow, _ := strconv.Atoi(matches[2])
	upperLeftColumn, _ := strconv.Atoi(matches[3])
	lowerRightRow, _ := strconv.Atoi(matches[4])
	lowerRightColumn, _ := strconv.Atoi(matches[5])

	through = CoordinatesRange{
		UpperLeft:  Coordinates{Row: upperLeftRow, Column: upperLeftColumn},
		LowerRight: Coordinates{Row: lowerRightRow, Column: lowerRightColumn},
	}

	return command, through, nil
}

func modifyLights(lights [][]int, command string, through CoordinatesRange) {
	for i := through.UpperLeft.Row; i <= through.LowerRight.Row; i++ {
		for j := through.UpperLeft.Column; j <= through.LowerRight.Column; j++ {
			switch command {
			case Toggle:
				lights[i][j] += 2
			case TurnOn:
				lights[i][j]++
			case TurnOff:
				if lights[i][j] > 0 {
					lights[i][j]--
				}
			default:
				panic("Doesn't exist!")
			}
		}
	}
}

func main() {
	N, M := 1000, 1000
	lights := make([][]int, N)
	for i := range lights {
		lights[i] = make([]int, M)
	}

	file, err := os.Open("instructions.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()

		command, through, err := parseInstruction(instruction)
		check(err)
		modifyLights(lights, command, through)
	}

	check(scanner.Err())

	brightness := 0

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			brightness += lights[i][j]
		}
	}

	fmt.Printf("Brightness: %v\n", brightness)
}
