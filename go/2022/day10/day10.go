package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type instructionName string

const (
	Addx instructionName = "addx"
	Noop instructionName = "noop"
)

type instruction struct {
	name   instructionName
	args   []string
	cycles int
}

func parseInstruction(input string) instruction {
	switch {
	case strings.Index(input, "noop") != -1:
		return instruction{
			name:   Noop,
			args:   []string{},
			cycles: 1,
		}
	case strings.Index(input, "addx") != -1:
		argIndex := strings.Index(input, " ") + 1
		return instruction{
			name:   Addx,
			args:   []string{input[argIndex:]},
			cycles: 2,
		}
	default:
		panic(fmt.Sprintf("invalid instruction: %v", input))
	}
}

func parse(input []string) []instruction {
	instructions := make([]instruction, 0)
	for _, inst := range input {
		instructions = append(instructions, parseInstruction(inst))
	}
	return instructions
}

func Part1(input []string) int {
	instructions := parse(input)
	currentInstruction := 0
	instructionCycleCount := 1
	xRegister := 1
	signalStrength := 0
	for c := 1; ; c++ {
		if len(instructions) <= currentInstruction {
			break
		}
		inst := instructions[currentInstruction]

		switch {
		case c == 20, (c+20)%40 == 0:
			signalStrength += c * xRegister
		}

		if instructionCycleCount == inst.cycles {
			currentInstruction++
			switch inst.name {
			case Noop:
				// noop
			case Addx:
				toAdd, _ := strconv.Atoi(inst.args[0])
				xRegister += toAdd
			}
			instructionCycleCount = 1
		} else {
			instructionCycleCount++
		}
	}

	return signalStrength
}
