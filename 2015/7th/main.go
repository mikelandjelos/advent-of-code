package main

import (
	"bufio"
	"fmt"
	"os"
)

// Operators.
const (
	Not    = "NOT"
	And    = "AND"
	Or     = "OR"
	LShift = "LSHIFT"
	RShift = "RSHIFT"
)

func emulateCircuit(circuit map[string]string) uint16 {

}

func createCircuit(instruction string) map[string]string {
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("assembly.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instruction := scanner.Text()
		fmt.Println(instruction)
		createCircuit(instruction)
		fmt.Println("=====================")
	}

	check(scanner.Err())
}
