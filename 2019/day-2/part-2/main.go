package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	inputs []int
)

func main() {
	// To do this, before running the program, replace position 1 with the value 12 and replace position 2 with the value 2
	inputs[1] = 12
	inputs[2] = 2

	answer := findAnswer()
	fmt.Printf("The answer is: %+v\n", int(answer))
}

func findAnswer() int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			resetInput()

			inputs[1] = noun
			inputs[2] = verb

			run(inputs)

			if inputs[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}

func run(inputs []int) ([]int, error) {
	for i := 0; i < len(inputs); i = i + 4 {
		opcode := inputs[i]
		if opcode == 99 {
			return inputs, nil
		}

		firstNumberPosition := inputs[i+1]
		secondNumberPosition := inputs[i+2]
		positionToStoreResult := inputs[i+3]
		if opcode == 1 {
			result := inputs[firstNumberPosition] + inputs[secondNumberPosition]
			inputs[positionToStoreResult] = result
		} else if opcode == 2 {
			result := inputs[firstNumberPosition] * inputs[secondNumberPosition]
			inputs[positionToStoreResult] = result
		} else {
			return nil, fmt.Errorf("unhandled case. Something went wrong")
		}
	}
	return inputs, nil
}

var inputByte []byte

func init() {
	resetInput()
}

func resetInput() {
	inputByte, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	inputArray := strings.FieldsFunc(string(inputByte), func(r rune) bool {
		return r == ','
	})

	inputs, err = parseStringArrayToInt(inputArray)
	if err != nil {
		panic(err)
	}
}

func parseStringArrayToInt(stringArray []string) ([]int, error) {
	intArray := make([]int, len(stringArray))
	var err error
	for i := 0; i < len(stringArray); i++ {
		intArray[i], err = strconv.Atoi(stringArray[i])
		if err != nil {
			return nil, err
		}
	}
	return intArray, nil
}
