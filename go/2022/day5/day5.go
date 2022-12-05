package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack[T any] []T

func Push[T any](slice []T, item T) []T {
	return append(slice, item)
}

func Pop[T any](slice []T) (T, []T) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func sampleStacks() []Stack[string] {
	return []Stack[string]{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
}

func inputStacks() []Stack[string] {
	return []Stack[string]{
		{"B", "W", "N"},
		{"L", "Z", "S", "P", "T", "D", "M", "B"},
		{"Q", "H", "Z", "W", "R"},
		{"W", "D", "V", "J", "Z", "R"},
		{"S", "H", "M", "B"},
		{"L", "G", "N", "J", "H", "V", "P", "B"},
		{"J", "Q", "Z", "F", "H", "D", "L", "S"},
		{"W", "S", "F", "J", "G", "Q", "B"},
		{"Z", "W", "M", "S", "C", "D", "J"},
	}
}

type instruction struct {
	from  int
	to    int
	count int
}

func parseStacks(s []string) []Stack[string] {
	return inputStacks()
}

func parseInstruction(s string) instruction {
	// move 3 from 5 to 2
	moveIdx := strings.Index(s, "move ")
	fromIdx := strings.Index(s, "from ")
	toIdx := strings.Index(s, "to ")
	move, _ := strconv.Atoi(s[moveIdx+5 : fromIdx-1])
	from, _ := strconv.Atoi(s[fromIdx+5 : toIdx-1])
	to, _ := strconv.Atoi(s[toIdx+3:])
	return instruction{
		from:  from,
		to:    to,
		count: move,
	}
}

func parseInput(input []string) (stacks []Stack[string], instructions []instruction) {
	emptyFound := false
	stackData := make([]string, 0)
	instructions = make([]instruction, 0)
	for _, line := range input {
		if line == "" {
			emptyFound = true
			continue
		}

		stackData = append(stackData, line)

		if emptyFound == false {
			continue
		}
		instructions = append(instructions, parseInstruction(line))
	}
	return parseStacks(stackData), instructions
}

func Part1(input []string) string {
	stacks, instructions := parseInput(input)

	for _, i := range instructions {
		for c := 0; c < i.count; c++ {
			item, newFrom := Pop(stacks[i.from-1])
			newTo := Push(stacks[i.to-1], item)
			stacks[i.from-1] = newFrom
			stacks[i.to-1] = newTo
		}
	}

	result := ""
	for _, stack := range stacks {
		pop, _ := Pop(stack)
		result += pop
	}
	return result
}

func Part2(input []string) string {
	stacks, instructions := parseInput(input)
	for _, i := range instructions {
		toMove := make([]string, 0)
		for c := 0; c < i.count; c++ {
			item, from := Pop(stacks[i.from-1])
			toMove = append(toMove, item)
			stacks[i.from-1] = from
		}

		Reverse(toMove)
		stacks[i.to-1] = append(stacks[i.to-1], toMove...)
	}

	result := ""
	for _, stack := range stacks {
		pop, _ := Pop(stack)
		result += pop
	}
	return result
	fmt.Printf("%v %v", stacks, instructions)
	return ""
}
