package day10

import (
	"aoc/util"
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

const (
	CRTWidth   int = 40
	CRTHeight  int = 6
	SpriteSize int = 3
)

func Part2(input []string) string {
	instructions := parse(input)
	currentInstruction := 0
	instructionCycleCount := 1
	xRegister := 1
	pixels := make([]string, 0)
	for c := 1; ; c++ {
		if len(instructions) <= currentInstruction {
			break
		}

		toDrawX := (c - 1) % 40
		toDrawY := int(float64(c-1) / float64(CRTWidth))
		if len(pixels) <= toDrawY {
			pixels = append(pixels, "")
		}
		pixel := "."
		if util.Abs(xRegister-toDrawX) <= 1 {
			pixel = "#"
		}

		pixels[toDrawY] += pixel

		inst := instructions[currentInstruction]
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
	draw(pixels)
	return ""
}

func draw(input []string) {
	fmt.Printf("\n\n")
	for i := 0; i < len(input); i++ {
		fmt.Printf("%v\n", input[i])
	}
	fmt.Printf("\n\n")
}
