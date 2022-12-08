package day5

import (
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

type instruction struct {
	from  int
	to    int
	count int
}

func parseStacks(s []string) []Stack[string] {
	stackCountRow := s[len(s)-1]
	numOfStacks, _ := strconv.Atoi(stackCountRow[len(stackCountRow)-1:])
	stackData := s[0 : len(s)-1]

	stacks := make([]Stack[string], numOfStacks)
	for _, row := range stackData {
		curStack := 0
		for i := 1; i < len(row); i += 4 {
			char := row[i : i+1]
			if char != " " {
				stacks[curStack] = append(stacks[curStack], char)
			}
			curStack++
		}
	}

	for _, stack := range stacks {
		Reverse(stack)
	}

	return stacks
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

	stackData := make([]string, 0)
	instructions = make([]instruction, 0)

	emptyFound := false
	for i := 0; i < len(input); i++ {
		line := input[i]
		if line == "" {
			emptyFound = true
			continue
		}

		if !emptyFound {
			stackData = append(stackData, line)
		}

		if emptyFound == true {
			instructions = append(instructions, parseInstruction(line))
		}
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
}
