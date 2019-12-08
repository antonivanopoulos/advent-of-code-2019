package main

import "log"

import "fmt"

// Intcode runs through the list of provided integers and performs the operations
// represented by those codes (addition, multiplication. and halting)
func Intcode(input int, program []int) []int {
	index := 0

	for index < len(program) {
		opcode := fmt.Sprintf("%05d", program[index])
		// paramMode3 := string(opcode[0])
		paramMode2 := string(opcode[1])
		paramMode1 := string(opcode[2])
		op := string(opcode[3:])

		switch op {
		case "01":
			var input1, input2 int

			if paramMode1 == "1" {
				input1 = program[index+1]
			} else {
				input1 = program[program[index+1]]
			}

			if paramMode2 == "1" {
				input2 = program[index+2]
			} else {
				input2 = program[program[index+2]]
			}

			program[program[index+3]] = input1 + input2
			index += 4
		case "02":
			var input1, input2 int

			if paramMode1 == "1" {
				input1 = program[index+1]
			} else {
				input1 = program[program[index+1]]
			}

			if paramMode2 == "1" {
				input2 = program[index+2]
			} else {
				input2 = program[program[index+2]]
			}

			program[program[index+3]] = input1 * input2
			index += 4
		case "03":
			program[program[index+1]] = input
			index += 2
		case "04":
			log.Printf("Output: %d\n", program[program[index+1]])
			index += 2
		case "99":
			return program
		}
	}

	return program
}
