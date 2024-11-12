package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/gookit/goutil/dump"
)

var logger = log.Default()

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Operators.
const (
	Not    = "NOT"
	And    = "AND"
	Or     = "OR"
	LShift = "LSHIFT"
	RShift = "RSHIFT"
)

type Signal = uint16

var Operators = map[string]func(...Signal) Signal{
	Not: func(signals ...Signal) Signal {
		return ^signals[0]
	},
	And: func(signals ...Signal) Signal {
		return signals[0] & signals[1]
	},
	Or: func(signals ...Signal) Signal {
		return signals[0] | signals[1]
	},
	LShift: func(signals ...Signal) Signal {
		return signals[0] << signals[1]
	},
	RShift: func(signals ...Signal) Signal {
		return signals[0] >> signals[1]
	},
}

var expressionCommonFormat = regexp.MustCompile(`^(?:(NOT) (\w+)|(\w+) (AND|OR|LSHIFT|RSHIFT) (\w+)|(\w+))`)
var assemblyInstructionCommonFormat = regexp.MustCompile(`(.*) -> (\w+)`)

func AssembleCircuit(instructionsFile string) map[string][]string {
	file, err := os.Open(instructionsFile)
	check(err)
	defer file.Close()

	circuit := make(map[string][]string)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		assemblyInstruction := scanner.Text()
		logger.Printf("Instruction: %v\n", assemblyInstruction)

		// Separating expression (left) and signal label (right) side.
		tokens := assemblyInstructionCommonFormat.FindStringSubmatch(assemblyInstruction)
		label := tokens[2]
		expression := tokens[1]

		// Tokenizing the expression.
		expressionTokens := expressionCommonFormat.FindStringSubmatch(expression)

		// Tokenize the expression.
		/* Explanation:
		=> expressionTokens[0] - full expression;
		=> expressionTokens[1:3] - reserved for `NOT <operand>` expression type;
		=> expressionTokens[3:6] - reserved for `<left-operand> <operator> <right-operand>` expression type;
		=> expressionTokens[6] - reserved for literals.
		*/
		switch {
		// Unary operator (NOT).
		case expressionTokens[1] != "" && expressionTokens[2] != "":
			circuit[label] = expressionTokens[1:3]

		// Binary operator (AND|OR|LSHIFT|RSHIFT).
		case expressionTokens[3] != "" && expressionTokens[4] != "" && expressionTokens[5] != "":
			circuit[label] = []string{expressionTokens[4], expressionTokens[3], expressionTokens[5]} // <operator> <left> <right>

			// No operators (simple assign).
		case expressionTokens[6] != "":
			circuit[label] = expressionTokens[6:]
		default:
			panic("Expression format not supported!")
		}
	}

	check(scanner.Err())

	return circuit
}

var literalCommonFormat = regexp.MustCompile(`(\d+)`)

func SimulateCircuit(circuit map[string][]string) map[string]uint16 {
	signals := make(map[string]uint16)

	lifo := NewStack[string]()

	for label := range circuit {
		if _, alreadyCalculated := signals[label]; alreadyCalculated {
			continue
		}

		// TODO: No way to find if there is a loop - this will just go endlessly if there is one.
		lifo.Push(label)
		for !lifo.Empty() {
			label, err := lifo.Pop()
			check(err)

			dump.P(signals)

			expressionTokens := circuit[label]

			switch len(expressionTokens) {
			case 1: // Value.
				labelOrLiteral := expressionTokens[0]

				if value, ok := signals[labelOrLiteral]; ok {
					// Label, already calculated.
					signals[label] = value
				} else if literalCommonFormat.MatchString(labelOrLiteral) {
					// Literal.
					value, err := strconv.Atoi(labelOrLiteral)
					check(err)
					signals[label] = Signal(value)
				} else {
					// Label, still not calculated.
					lifo.Peek(label)
					lifo.Push(labelOrLiteral)
				}
			case 2: // Unary expression.
				operator, operand := expressionTokens[0], expressionTokens[1]
				if value, ok := signals[operand]; ok {
					signals[label] = Operators[operator](value)
				} else if literalCommonFormat.MatchString(operand) {
					// Literal.
					value, err := strconv.Atoi(operand)
					check(err)
					signals[label] = Signal(value)
				} else {
					lifo.Push(label)
					lifo.Push(operand)
				}
			case 3: // Binary expression.
				operator, leftOperand, rightOperand := expressionTokens[0], expressionTokens[1], expressionTokens[2]
			}
		}
	}

	dump.P(signals)

	panic("Not implemented yet.")

	return signals
}

func main() {
	circuit := AssembleCircuit("assembly.txt")
	results := SimulateCircuit(circuit)

	resultLabel := "i"
	resultValue, exists := results[resultLabel]

	if !exists {
		panic(fmt.Errorf("Signal `%v` doesn't exist!", resultLabel))
	}

	logger.Printf("`%v` -> %v", resultLabel, resultValue)
}
