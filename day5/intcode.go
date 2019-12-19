package main

import (
	"log"
)

type Function int

const (
	Addition       Function = 1
	Multiplication Function = 2
	Input          Function = 3
	Output         Function = 4
	JumpIfTrue     Function = 5
	JumpIfFalse    Function = 6
	LessThan       Function = 7
	Equals         Function = 8
	Halt           Function = 99
)

type Operation struct {
	function Function
	index    int
	p1mode   int
	p2mode   int
	p3mode   int
}

// Intcode runs through the list of provided integers and performs the operations
// represented by those codes (addition, multiplication. and halting)
func Intcode(input int, program []int) []int {
	index := 0

	for index < len(program) {
		operation := ParseOpCode(program[index])
		operation.index = index

		switch operation.function {
		case Addition:
			input1 := ParseInput(program, operation.index+1, operation.p1mode)
			input2 := ParseInput(program, operation.index+2, operation.p2mode)
			input3 := program[operation.index+3]

			program[input3] = input1 + input2
			index += 4
		case Multiplication:
			input1 := ParseInput(program, operation.index+1, operation.p1mode)
			input2 := ParseInput(program, operation.index+2, operation.p2mode)
			input3 := program[operation.index+3]

			program[input3] = input1 * input2
			index += 4
		case Input:
			program[program[index+1]] = input
			index += 2
		case Output:
			input1 := ParseInput(program, operation.index+1, operation.p1mode)
			log.Printf("Output: %d\n", input1)
			index += 2
		case JumpIfTrue:
			input1 := ParseInput(program, operation.index+1, operation.p1mode)
			input2 := ParseInput(program, operation.index+2, operation.p2mode)

			if input1 != 0 {
				index = input2
			} else {
				index += 3
			}
		case JumpIfFalse:
			input1 := ParseInput(program, operation.index+1, operation.p1mode)
			input2 := ParseInput(program, operation.index+2, operation.p2mode)

			if input1 == 0 {
				index = input2
			} else {
				index += 3
			}
		case LessThan:
			input1 := ParseInput(program, operation.index+1, operation.p1mode)
			input2 := ParseInput(program, operation.index+2, operation.p2mode)
			input3 := program[index+3]

			if input1 < input2 {
				program[input3] = 1
			} else {
				program[input3] = 0
			}

			index += 4
		case Equals:
			input1 := ParseInput(program, operation.index+1, operation.p1mode)
			input2 := ParseInput(program, operation.index+2, operation.p2mode)
			input3 := program[index+3]

			if input1 == input2 {
				program[input3] = 1
			} else {
				program[input3] = 0
			}

			index += 4
		case Halt:
			return program
		default:
			panic("unknown opcode")
		}
	}

	return program
}

func ParseOpCode(input int) Operation {
	opcode := input % 100
	modes := input / 100

	p1mode := modes % 10
	p2mode := (modes / 10) % 10
	p3mode := (modes / 100) % 10
	var op Function = Function(opcode)

	return Operation{op, 0, p1mode, p2mode, p3mode}
}

func ParseInput(program []int, index int, opmode int) int {
	if opmode == 1 {
		return program[index]
	}

	return program[program[index]]
}
